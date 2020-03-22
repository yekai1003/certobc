pragma solidity^0.6.1;
//pragma experimental ABIEncoderV2;

contract cert {
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
    function getHash() public view returns (bytes32) {
        return keccak256(abi.encode(msg.sender));
    }
     // 证书发行:输入用户uuid和课程名称hash值
    function issue(string memory uuid, bytes32 olHash) public onlyadmin {
        bytes32 hash = keccak256(abi.encode(uuid, olHash));
        require(!certMap[hash].ok, "must not exists");
        Cert memory cert1 = Cert(uuid, true, olHash);
        certMap[hash] = cert1;
        userCerts[uuid].push(hash);
        emit certIssue(uuid, olHash, hash);
    }
    // 证书验证:输入用户uuid和课程名称hash值
    function verifyCourseHash(string memory uuid, bytes32 olHash) public view returns (bytes32) {
        bytes32 hash = keccak256(abi.encode(uuid, olHash));
        require(certMap[hash].ok, "must be exists");
        require(certMap[hash].olHash ==olHash, "hash must right!");
        return hash;
    }
    function getUserCerts(string memory uuid) public view returns (bytes32[] memory) {
        return userCerts[uuid];
    }
    
    
    
}