# Decentralized Voting System

This project is a decentralized voting system built using **Go**, **Ethereum Smart Contracts**, and **Gin Web Framework**. It allows users to register, log in, and vote for candidates in a decentralized manner. The votes are stored on the Ethereum blockchain, ensuring transparency and immutability.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Smart Contract](#smart-contract)
- [API Endpoints](#api-endpoints)
- [Setup and Installation](#setup-and-installation)
- [Environment Variables](#environment-variables)
- [How It Works](#how-it-works)
- [Testing](#testing)
- [Future Improvements](#future-improvements)

---

## 🔍 Overview

The decentralized voting system leverages blockchain technology to ensure secure and transparent voting. The backend is implemented in Go, and the voting logic is handled by an Ethereum smart contract deployed on a local Ganache blockchain. The system includes features like user registration, login with JWT authentication, and vote submission with transaction tracking.

---

## 🛠️ Technologies Used

### Backend
- **Go 1.24.3**: The primary programming language for the backend
- **Gin**: A lightweight web framework for building RESTful APIs
- **GORM**: An ORM library for database interactions with MySQL
- **MySQL**: The database used for storing user information
- **JWT (golang-jwt/jwt)**: For secure user authentication and session management

### Blockchain
- **Solidity ^0.8.0**: Used to write the Ethereum smart contract
- **Ganache**: A local Ethereum blockchain for development and testing (running on port 7545)
- **Remix IDE**: Used to compile and deploy the smart contract
- **go-ethereum**: A Go library for interacting with Ethereum blockchain
- **abigen**: Used to generate Go bindings from smart contract ABI

### Security & Utilities
- **MD5 Hashing**: For password security (stored in `utils/security/hash.go`)
- **godotenv**: For managing environment variables
- **charmbracelet/log**: For structured and colorful logging
- **sethvargo/go-envconfig**: For environment configuration management

### Development Tools
- **Git**: Version control (with `.gitignore` for sensitive files)
- **Environment Configuration**: Using `.env` files for local development

---

## 📁 Project Structure

```
├── .env                    # Environment variables (gitignored)
├── .env.example           # Environment variables template
├── .gitignore             # Git ignore rules
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
├── main.go                # Application entry point
├── README.md              # Project documentation
├── config/
│   └── config.go          # Configuration management
├── contracts/
│   ├── voting.go          # Generated Go bindings for smart contract
│   ├── abi/
│   │   └── voting.abi.json    # Smart contract ABI
│   └── sol/
│       └── voting.sol     # Solidity smart contract
├── controller/
│   ├── health.go          # Health check endpoint
│   ├── users.go           # User registration and authentication
│   └── vote.go            # Voting functionality
├── middleware/
│   ├── jwt/
│   │   └── jwt.go         # JWT authentication middleware
│   └── logger/
│       └── logger.go      # Custom logging middleware
├── models/
│   ├── init.go            # Database initialization
│   └── models.go          # Database models (User)
├── router/
│   └── http/
│       ├── init.go        # Router initialization
│       └── router.go      # Route definitions
├── service/
│   └── contract_service.go    # Blockchain interaction service
└── utils/
    ├── jwt/
    │   └── jwt.go         # JWT token utilities
    └── security/
        └── hash.go        # Password hashing utilities
```

---

## 📜 Smart Contract

### Contract Details
- **File**: `contracts/sol/voting.sol`
- **Language**: Solidity ^0.8.0
- **License**: MIT
- **Deployed Address**: `0x195d78AB4ECDC9C8FB2473aaeA9BEfe3832C3340`

### Contract Functions
```solidity
contract Voting {
    mapping(string => uint) public votes;

    function vote(string memory candidate) public {
        votes[candidate]++;
    }

    function getVotes(string memory candidate) public view returns (uint) {
        return votes[candidate];
    }
}
```

### Features
- **Vote Storage**: Mapping of candidate names to vote counts
- **Vote Function**: Allows voting for any candidate (string-based)
- **Vote Retrieval**: Public function to get current vote count for any candidate
- **Gas Efficient**: Simple increment operation for voting

---

## 🌐 API Endpoints

### Base URL
```
http://localhost:2020/v1
```

### Public Endpoints

#### Health Check
```http
GET /v1/health
```
**Response:**
```json
{
  "status": "ok",
  "message": "Service is running."
}
```

#### User Registration
```http
POST /v1/register
```
**Request Body:**
```json
{
  "username": "john_doe",
  "password": "secure_password"
}
```
**Response:**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "username": "john_doe"
  }
}
```

#### User Login
```http
POST /v1/login
```
**Request Body:**
```json
{
  "username": "john_doe",
  "password": "secure_password"
}
```
**Response:**
```json
{
  "message": "Login successful",
  "user": {
    "id": 1,
    "username": "john_doe",
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### Get Vote Count
```http
GET /v1/votes?candidate=john_smith
```
**Response:**
```json
{
  "candidate": "john_smith",
  "votes": "42"
}
```

### Protected Endpoints (Requires JWT Token)

#### Submit Vote
```http
POST /v1/vote
Authorization: Bearer <jwt_token>
```
**Request Body:**
```json
{
  "candidate": "john_smith"
}
```
**Response:**
```json
{
  "message": "Vote submitted successfully",
  "tx_hash": "0x1234567890abcdef..."
}
```

---

## ⚙️ Setup and Installation

### Prerequisites
- **Go 1.24.3 or higher**
- **MySQL Server**
- **Ganache CLI or Ganache GUI**
- **Git**

### 1. Clone the Repository
```bash
git clone <your-repository-url>
cd Voting-system
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Setup Ganache
```bash
# Install Ganache CLI globally
npm install -g ganache-cli

# Start Ganache on port 7545
ganache-cli -p 7545 -d
```

### 4. Deploy Smart Contract
1. Open **Remix IDE** (https://remix.ethereum.org/)
2. Create a new file and paste the content from `contracts/sol/voting.sol`
3. Compile the contract using Solidity compiler ^0.8.0
4. Connect to your local Ganache network (HTTP://127.0.0.1:7545)
5. Deploy the contract and note the contract address
6. Copy the ABI and save it to `contracts/abi/voting.abi.json`

### 5. Generate Go Bindings (Optional)
```bash
# Install abigen if not already installed
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Generate Go bindings from ABI
abigen --abi=contracts/abi/voting.abi.json --pkg=contract --out=contracts/voting.go
```

### 6. Setup MySQL Database
```sql
CREATE DATABASE voting_system;
```

### 7. Configure Environment Variables
```bash
cp .env.example .env
```
Edit `.env` with your configuration:
```env
PORT=2020
HOST=localhost
RPC_URL=http://127.0.0.1:7545
CONTRACT_ADDRESS=<your_deployed_contract_address>
PRIVATE_KEY=<your_ganache_private_key>
DB_User=root
DB_Passwd=<your_mysql_password>
DB_Host=localhost
DB_Port=3306
DB_Name=voting_system
JWT_SECRET=<your_jwt_secret>
```

### 8. Run the Application
```bash
go run main.go
```

The server will start on `http://localhost:2020`

---

## 🔧 Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | Server port | 8080 | No |
| `HOST` | Server host | localhost | No |
| `RPC_URL` | Ethereum RPC URL | http://localhost:8545 | No |
| `CONTRACT_ADDRESS` | Deployed contract address | - | Yes |
| `PRIVATE_KEY` | Ethereum private key for transactions | - | Yes |
| `DB_User` | MySQL username | - | Yes |
| `DB_Passwd` | MySQL password | - | Yes |
| `DB_Host` | MySQL host | - | Yes |
| `DB_Port` | MySQL port | 3306 | No |
| `DB_Name` | MySQL database name | - | Yes |
| `JWT_SECRET` | JWT signing secret | test | No |

---

## 🚀 How It Works

### 1. User Registration & Authentication
- Users register with username/password
- Passwords are hashed using MD5 (stored in database)
- JWT tokens are issued upon successful login
- Tokens expire after 24 hours

### 2. Voting Process
- Authenticated users can submit votes via API
- Each vote triggers a blockchain transaction
- The smart contract increments the vote count for the specified candidate
- Transaction hash is returned to the user for verification

### 3. Vote Retrieval
- Anyone can query vote counts for any candidate
- Data is retrieved directly from the blockchain
- No authentication required for reading vote counts

### 4. Blockchain Integration
- Uses go-ethereum library for blockchain interaction
- Connects to local Ganache network for development
- Automatic gas price estimation and nonce management
- Transaction signing with private key

---

## 🧪 Testing

### Manual Testing with cURL

#### Register a User
```bash
curl -X POST http://localhost:2020/v1/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

#### Login
```bash
curl -X POST http://localhost:2020/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"testpass"}'
```

#### Submit a Vote
```bash
curl -X POST http://localhost:2020/v1/vote \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
  -d '{"candidate":"alice"}'
```

#### Check Vote Count
```bash
curl "http://localhost:2020/v1/votes?candidate=alice"
```

### Database Testing
- User data is stored in MySQL with GORM auto-migration
- Passwords are hashed before storage
- Soft deletes are supported with GORM

### Blockchain Testing
- Test on local Ganache network
- Monitor transactions in Ganache GUI
- Verify vote counts directly on blockchain

---

## 🔮 Future Improvements

### Security Enhancements
- [ ] Replace MD5 with bcrypt for password hashing
- [ ] Implement rate limiting for API endpoints
- [ ] Add input validation and sanitization
- [ ] Implement CORS middleware
- [ ] Add HTTPS support

### Blockchain Features
- [ ] Add candidate registration functionality
- [ ] Implement voting time windows
- [ ] Add vote verification mechanisms
- [ ] Support for multiple concurrent elections
- [ ] Event logging for audit trails

### API Improvements
- [ ] Add comprehensive error handling
- [ ] Implement API versioning
- [ ] Add Swagger/OpenAPI documentation
- [ ] Add pagination for large datasets
- [ ] Implement WebSocket for real-time updates

### Infrastructure
- [ ] Add Docker containerization
- [ ] Implement CI/CD pipeline
- [ ] Add comprehensive unit tests
- [ ] Setup monitoring and metrics
- [ ] Add database migrations

### Frontend Integration
- [ ] Build React/Vue.js frontend
- [ ] Add MetaMask integration
- [ ] Real-time vote count updates
- [ ] Responsive mobile design

---

## 📝 License

This project is licensed under the MIT License - see the smart contract header for details.

---

## 👨‍💻 Author

**emenesism**
- GitHub: [@emenesism](https://github.com/emenesism)

---

## 🙏 Acknowledgments

- **Ethereum Foundation** for go-ethereum library
- **Gin Framework** for the lightweight web framework
- **GORM** for excellent ORM capabilities
- **Ganache** for local blockchain development
- **Remix IDE** for smart contract development and deployment

---

## 📞 Support

If you encounter any issues or have questions, please:
1. Check the logs for error messages
2. Ensure all dependencies are properly installed
3. Verify your environment variables are correctly set
4. Make sure Ganache is running and accessible
5. Check that the smart contract is deployed correctly

For additional help, please open an issue in the repository.
