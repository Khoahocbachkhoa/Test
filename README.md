# TEST

Học Devops

## Kiến trúc

## Backend

Backend đơn giản viết bằng Go và tương tác với cơ sỡ dữ liệu psql.

Khởi tạo kết nối tới cơ sỡ dữ liệu postgres, và chạy ứng dụng backend:

```
go mod download
go run .
```

## Docker compose

Cần đóng gói 3 services bằng docker : Go app, database và Nginx

Tạo docker volume cho database để lưu dữ liệu bền vững 

Truyền biến môi trường có chuỗi kết nối cơ sỡ dữ liệu cho container của backend app.

Cần thực hiện healthcheck cho database để đảm bảo database được khởi tạo đúng và Go App kết nối thành công

Truyền config Nginx đóng vai trò reverse Proxy và làm entry point cho toàn bộ ứng dụng.

Đóng gói ứng dụng bằng docker compose và chạy:

```
docker compose up -d
```

Check service, logs:

```
docker compose ps
docker compose logs nginx
docker compose logs app
docker compose logs db
```

Test:

```
curl http://localhost/health
```

## Linux Server Deployment

Tạo một máy ảo Linux và triển khai lên đó.

Có thể dùng một số phần mềm như VMware, Virtual Box hoặc KVM/Qemu.

Tạo một máy ảo Ubuntu server tối thiểu 2Gb RAM, 20Gb Storage.

Thực hiện ssh vào server và chạy:

```
git clone https://github.com/Khoahocbachkhoa/Test
cd Test
docker compose up -d
```

Truy cập từ host vào:

```
curl http://<ip của máy ảo kết nối với host>/health
```

## Kubernetes

Sử dụng k3s (Lightweight Kubernetes) để triển khai kubernetes trên VM Linux Machine.

Tạo thêm 2 máy ảo Linux làm 2 worker node tương ứng với master node là máy ảo ở trên.

Kiến trúc:

```

                    Host
                         │
                  libvirt / NAT
                         │
              192.168.122.0/24
                         │
        ┌────────────────┼────────────────┐
        │                │                │
        ▼                ▼                ▼
   VM1 - Master      VM2 - Worker 1   VM3 - Worker 2
   192.168.122.10    192.168.122.11   192.168.122.12
        │                │                │
        └────────────────┼────────────────┘
                         │
                  Internal network
```