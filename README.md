# TEST

Học devops

## Kiến trúc

## Backend

Khởi tạo cơ sỡ dữ liệu postgres, tạo biến môi trường và chạy:

```
go mod download
go run .
```

## Docker compose

Đóng gói ứng dụng bằng docker compose và chạy:

```
docker compose up -d
```

Check service, logs:

```
docker compose ps
docker compose logs
```

Test:

```
curl http://localhost/health
```

## Linux Server Deployment

Tạo một máy ảo Linux và triển khai lên đó.
Có thể dùng một số phần mềm như VMware, Virtual Box hoặc KVM/Qemu.
Thực hiện ssh vào server đó và chạy:

```
git clone https://github.com/Khoahocbachkhoa/Test
cd Test
docker compose up -d
```

## Kubernetes

