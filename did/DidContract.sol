pragma solidity ^0.6.10;

import "./LibBytesMap.sol";
import "@openzeppelin/contracts@3.4.0/access/Ownable.sol";

/**
 * @title DidContract
 */
contract DidContract is Ownable {
    using LibBytesMap for LibBytesMap.Map;
    LibBytesMap.Map private didMap;
    mapping(bytes => address) internal Possessor;

    /**
     * @dev Error code
     */
    string private constant SUCCESS = "0000";
    string private constant DID_NOT_EXIST = "2011";
    string private constant DID_ALREADY_EXIST = "2012";

    event createDidRetLog(string msgcode, string msg);

    event updateDocumentRetLog(string msgcode, string msg);

    event getInitializeDataLog(string msgcode, bytes msg);

    constructor() public {}


    /**
     * @dev Create did
     *
     * @param did the did
     * @param didDocument the didDocument
     * @return msgcode the code for result
     * @return msg the msg for result
     */
    function createDid(string memory did, string memory didDocument)
        public
        returns (string memory msgcode, string memory msg)
    {
        bool bools = isDidExist(did);
        if (bools) {
            emit createDidRetLog(DID_ALREADY_EXIST, "did already exist");
            return (DID_ALREADY_EXIST, "did already exist");
        }
        bytes memory didBytes = bytes(did);
        bytes memory didDocumentBytes = bytes(didDocument);
        didMap.put(didBytes, didDocumentBytes);
        Possessor[didBytes] = _msgSender();
        emit createDidRetLog(SUCCESS, "success");
        return (SUCCESS, "success");
    }

    /**
     * @dev Get did document
     *
     * @param did the did
     * @return msgcode the code for result
     * @return msg the msg for result
     * @return didDocument the didDocument for result
     */
    function getDocument(string memory did)
        public
        view
        returns (
            string memory msgcode,
            string memory msg,
            string memory didDocument
        )
    {
        bool bools = isDidExist(did);
        if (!bools) {
            return (DID_NOT_EXIST, "did not exist", "");
        }
        bytes memory didBytes = bytes(did);
        bytes memory didDocumentBytes = didMap.getValue(didBytes);
        return (SUCCESS, "success", string(didDocumentBytes));
    }

    /**
     * @dev Update did document
     *
     * @param did the did
     * @param didDocument the didDocument
     * @return msgcode the code for result
     * @return msg the msg for result
     */
    function updateDocument(string memory did, string memory didDocument)
        public
        returns (string memory msgcode, string memory msg)
    {
        bytes memory didBytes = bytes(did);
        bytes memory didDocumentBytes = bytes(didDocument);
        require(
            Possessor[didBytes] == _msgSender(),
            "DidContract: caller is not the Possessor"
        );
        bool bools = isDidExist(did);
        if (!bools) {
            emit updateDocumentRetLog(DID_NOT_EXIST, "did not exist");
            return (DID_NOT_EXIST, "did not exist");
        }

        didMap.put(didBytes, didDocumentBytes);
        emit updateDocumentRetLog(SUCCESS, "success");
        return (SUCCESS, "success");
    }

    /**
     * @dev Judge whether did exist, return true if it exist
     *
     * @param did the did
     * @return bool the bool for result
     */
    function isDidExist(string memory did) public view returns (bool) {
        bytes memory key = bytes(did);
        uint256 idx = didMap.index[key];
        if (idx == 0) {
            return false;
        }
        return true;
    }
}
