# Upvote
Technical Challenger Klever.io


: Coin Service
  - gRPC calls using grpcurl
    -> Create coin
      # grpcurl --plaintext -d '{ "Description": "COIN 1", "Short": "C1" }' localhost:9090 CoinService/CreateCoin

    -> Delete coin by ID (soft delete)
      # grpcurl --plaintext -d '{ "ID": 1 }' localhost:9090 CoinService/DeleteCoin