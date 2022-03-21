cd ~/queojv4/

kill $(ps -ef | grep ./record | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./user | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./ask | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./message | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./problem | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./solution | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./email | grep -v grep |awk '{print $2}')
kill $(ps -ef | grep ./queoj | grep -v grep |awk '{print $2}')

go build service/user/user.go
go build service/ask/ask.go
go build service/solution/solution.go
go build service/email/email.go
go build service/message/message.go
go build service/problem/problem.go
go build service/record/record.go
go build bff/queoj.go

nohup ./email  > ./logs/email.log 2>&1 &
sleep 1

nohup ./ask  > ./logs/ask.log 2>&1 &
sleep 1

nohup ./solution  > ./logs/solution.log 2>&1 &
sleep 1

nohup ./message  > ./logs/message.log 2>&1 &
sleep 1

nohup ./problem  > ./logs/problem.log 2>&1 &
sleep 1

nohup ./user  > ./logs/user.log 2>&1 &
sleep 1

nohup ./record  > ./logs/record.log 2>&1 &
sleep 1

nohup ./queoj  > ./logs/queoj.log 2>&1 &


