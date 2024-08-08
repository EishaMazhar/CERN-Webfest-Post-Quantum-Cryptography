## WEBfest Hackathon Challenge: 

#### Post Quantum Cryptography: Integrating PQDS into Web3 with COSMOS SDK

## What is Post Quantum cryptography?
Post-quantum cryptography (PQC), sometimes referred to as quantum-proof, quantum-safe, or quantum-resistant, is the development of cryptographic algorithms (usually public-key algorithms) that are thought to be secure against a cryptanalytic attack by a quantum computer. Most widely-used public-key algorithms rely on the difficulty of one of three mathematical problems: the integer factorization problem, the discrete logarithm problem or the elliptic-curve discrete logarithm problem. (Taken from Wikipedia)

## PQDS algorithms Algorithm

The National Institute of Standards and Technology (NIST) has standardized three post-quantum digital signature (PQDS) algorithms:

1. **CRYSTALS-DILITHIUM**  
2. **FALCON**  
3. **SPHINCS+**

| Features | CRYSTALS-DILITHIUM | FALCON | SPHINCS+ |
| :---- | :---- | :---- | :---- |
| **Security Basis** | Lattice-based (Learning With Errors) | Lattice-based (NTRU, Fast Fourier Transforms) | Hash-based (Merkle Trees, Hash Functions) |
| **Signature Size** | Moderate | Smallest | Largest |
| **Key Size** | Moderate | Small | Largest |
| **Performance** | Balanced (moderate speed and efficiency) | Fast (efficient signing and verification) | Slow (due to extensive hash computations) |
| **Implementation Complexity** | Relatively simple | High (complex FFTs) | Moderate |
| **Stateless** | No | No | Yes |
| **Memory Usage** | Moderate | High (due to FFTs) | Low |
| **Computational Overhead** | Moderate | Low | High |
| **Suitability for Blockchain** | A good balance for various applications | Best for minimizing data size | Best for stateless key management |
| **Challenges** | Larger signatures, and optimization needed | Complex implementation, resource demands | Large signatures, slow performance |
|  |  |  |  |

## Use Cases

For implementation on the COSMOS SDK, the use cases need to be feasible within the blockchain framework and leverage the modular architecture and interoperability features of the COSMOS SDK. Here are three appropriate use cases that could be effectively implemented using the COSMOS SDK:

### **1\. Decentralized Identity (DID) for Researchers**

**Description:** Develop a decentralized identity system\[1\] for researchers that uses PQDS to secure identity verification and credential management.

**Algorithm Choice:** CRYSTALS-DILITHIUM for balanced performance and security.

**Features:**

* **Creation and Management of Researcher Identities:** Researchers can create and manage their identities on the blockchain.  
* **Secure Credential Issuance and Verification:** Institutions can issue and verify credentials using PQDS, ensuring authenticity and integrity.  
* **Integration with Academic and Research Institutions:** Allow institutions to validate credentials and manage researcher profiles.

**Implementation on COSMOS SDK:**

* **Identity Module:** Develop a module to handle identity creation, credential issuance, and verification.  
* **PQDS Integration:** Implement CRYSTALS-DILITHIUM for signing and verifying credentials.  
* **Interoperability:** Leverage COSMOS SDK’s IBC (Inter-Blockchain Communication) to integrate with other blockchains and institutional databases.

### **2\. Quantum-Resistant Voting System for Collaborative Projects**

**Description:** Build a secure, quantum-resistant voting system for decision-making in collaborative scientific projects and DAOs.

**Algorithm Choice:** CRYSTALS-DILITHIUM or FALCON for efficient and secure voting.

**Features:**

* **Secure Proposal Submission and Voting:** Enable secure submission of proposals and casting of votes using PQDS.  
* **Transparent and Verifiable Voting Results:** Ensure that all votes are transparently recorded and verifiable.  
* **Privacy and Integrity of Votes:** Use PQDS to maintain the integrity and privacy of votes.

**Implementation on COSMOS SDK:**

* **Voting Module:** Create a module for proposal submission, vote casting, and tallying.  
* **PQDS Integration:** Use CRYSTALS-DILITHIUM or FALCON for signing and verifying votes.  
* **Governance Integration:** Integrate with existing governance modules in the COSMOS SDK for seamless operation within DAOs.

### **3\. Supply Chain Management for Scientific Equipment**

**Description:** Implement a blockchain-based supply chain management system for tracking and verifying the authenticity of scientific equipment.

**Algorithm Choice:** CRYSTALS-DILITHIUM or FALCON for efficient tracking.

**Features:**

* **Secure Logging of Equipment Manufacturing, Shipping, and Receipt:** Track the equipment lifecycle from manufacture to delivery.  
* **Verification of Equipment Authenticity:** Use PQDS to ensure the authenticity of equipment at each stage.  
* **Transparent and Immutable Supply Chain Records:** Maintain transparent and immutable records on the blockchain.

**Implementation on COSMOS SDK:**

* **Supply Chain Module:** Develop a module for tracking and logging equipment status at various stages.  
* **PQDS Integration:** Implement CRYSTALS-DILITHIUM or FALCON for signing and verifying transaction logs.  
* **Interoperability:** Use IBC to facilitate cross-chain data sharing and integration with external systems.

## Conclusion

In summary:

* **CRYSTALS-DILITHIUM** offers a good balance and will likely be the most versatile for various blockchain applications.  
* **FALCON** is best for scenarios where data size needs to be minimized, but its complexity could be a hurdle.  
* **SPHINCS+** provides robust security and simplifies key management, but its large signature size and slower performance may limit its use in high-frequency transaction environments.

## 

## Reference:

Cosmos SDK Tutorial: [https://github.com/cosmos/sdk-tutorials](https://github.com/cosmos/sdk-tutorials)  
[https://www.mdpi.com/1999-4893/16/11/518\#:\~:text=The%20two%20lattice%2Dbased%20algorithms,has%20the%20largest%20signature%20length](https://www.mdpi.com/1999-4893/16/11/518\#:\~:text=The%20two%20lattice%2Dbased%20algorithms,has%20the%20largest%20signature%20length).  
\[1\] [https://www.researchgate.net/publication/339835028\_Decentralized\_Identity\_Where\_Did\_It\_Come\_From\_and\_Where\_Is\_It\_Going](https://www.researchgate.net/publication/339835028\_Decentralized\_Identity\_Where\_Did\_It\_Come\_From\_and\_Where\_Is\_It\_Going)

