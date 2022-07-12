# Upvote
Technical Challenger Klever.io


# Coin Service
  - gRPC calls using grpcurl
    - Create coin
      # grpcurl --plaintext -d '{ "Description": "COIN 1", "Short": "C1" }' localhost:9090 CoinService/CreateCoin

    - Delete coin by ID (soft delete)
      # grpcurl --plaintext -d '{ "ID": 1 }' localhost:9090 CoinService/DeleteCoin

    - Update coin
      # grpcurl --plaintext -d '{ "ID": 1, "Description": "New Description", "Short": "NewShort" }' localhost:9090 CoinService/UpdateCoin

    - List active coins
      # grpcurl --plaintext localhost:9090 CoinService/ListActiveCoins

    - Coin vote up
      # grpcurl --plaintext -d '{ "CoinID": 1 }' localhost:9090 CoinService/VoteUP

    - Coin vote down
      # grpcurl --plaintext -d '{ "CoinID": 1 }' localhost:9090 CoinService/VoteDown

# Monitor Coin Service
  - gRPC calls using grpcurl
    - # grpcurl --plaintext -d '{ "CoinID": 2 }' localhost:9090 VoteMonitor/CoinMonitorVotes