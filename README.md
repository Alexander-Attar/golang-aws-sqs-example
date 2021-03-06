# SQS 需要 IAM 管理的 Access Key 及 Secret Key

SQS 這兩把 key 是使用 IAM 管理, 所以在 IAM 那點擊 User 會看到 Access Key,

如果當初管理者沒有把 Secret 記下來, 那麼就要在該頁方點擊 `Manage Access Keys` 再點 `Create Access Key` 重新建立,

才會顯示 Secret Key, 注意! 記得在關掉 popup 前把 Secret Key 抄下來

**要記得把 SQS 的權限給你的 User, 點擊 `Attach User Policy`**

# Install

    go get github.com/crowdmob/goamz/sqs
    go get github.com/crowdmob/goamz/aws

# Type your access key, secret key, SQS queue name and region (enqueue.go and main.go)

    accessKey = "*************"
    secretKey = "*************"
    queueName = https://sqs.ap-northeast-1.amazonaws.com/5**********5/TestQueue

    mySqs := sqs.New(auth, aws.APNortheast)  <= Change to your current region in enqueue.go and main.go

> [region list](https://github.com/crowdmob/goamz/blob/master/aws/regions.go)

# Run enqueue

    mv enqueue.go /tmp/
    go run /tmp/enqueue.go

# Run worker

    go build
    ./golang-aws-sqs-example


# You will see (in worker) :

    2014/10/29 12:22:22 worker: Start polling
    2014/10/29 12:22:22 worker: Received 1 messages
    2014/10/29 12:22:22 worker: Spawned worker goroutine
    This is a test message from golang
    2014/10/29 12:22:22 worker: Start polling
    2014/10/29 12:22:22 worker: Start polling

# Multi-worker

利用 redis 做每個 job 的 flag, 所以可在 lan 環境下在不同主機開多個 worker 去跑

# Others

每個 SQS job 都有個 timeout, 如果撈出來後一段時間(約 30 秒)沒被刪掉, 再撈可能還會再撈到它, 建議做個
flag, 確保 job 沒被重複執行

# 參考及修改來源 `https://github.com/nabeken/golang-sqs-worker-example`
