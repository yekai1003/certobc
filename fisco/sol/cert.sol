pragma solidity^0.4.25;
pragma experimental ABIEncoderV2;

contract olcert {
    address public admin;
    struct Cert {
        string uuid;
        bool ok;
        bytes32 olHash;//outline-hash
    }
    //cert-hash => Cert 
    mapping(bytes32=>Cert) certMap;
    mapping(string=>bytes32[]) userCerts; 
    event certIssue(string  uuid, bytes32 olhash, bytes32 certHash);
    modifier onlyadmin()  {
        require(msg.sender == admin, "only admin can do");
        _;
    }
    
    constructor() public {
        admin = msg.sender;
    }
    function setAdmin(address user) public onlyadmin {
        require(address(0) != user, "user must exists");
        admin = user;
    }
    function issue(string memory uuid, bytes32 olHash) public onlyadmin {
        bytes32 hash = keccak256(abi.encode(uuid, olHash));
        require(!certMap[hash].ok, "must not exists");
        Cert memory cert1 = Cert(uuid, true, olHash);
        certMap[hash] = cert1;
        userCerts[uuid].push(hash);
        emit certIssue(uuid, olHash, hash);
    }
    function getCourseByHash(bytes32 certHash) public view returns (Cert memory) {
        return certMap[certHash];
    }
    function getUserCerts(string memory uuid) public view returns (bytes32[] memory) {
        return userCerts[uuid];
    }
    function verifyCourseHash(string memory uuid, bytes32 olHash) public view returns (bytes32) {
        bytes32 hash = keccak256(abi.encode(uuid, olHash));
        require(certMap[hash].ok, "must be exists");
        require(certMap[hash].olHash ==olHash, "hash must right!");
        return hash;
    }
    
    
}