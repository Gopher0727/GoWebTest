package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"

	"github.com/Gopher0727/GoWebTest/model"
)

func main() {
	data, err := os.ReadFile("model.toml")
	if err != nil {
		log.Fatalln(err)
	}

	var cfg model.Config
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	os.MkdirAll("data/mysql_data", 0755)
	os.MkdirAll("data/redis_data", 0755)

	err = os.WriteFile("docker-compose.yml", []byte(generateDockerCompose(cfg)), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("docker-compose.yml 已生成成功！执行 docker compose pull 即可")
}

func generateDockerCompose(cfg model.Config) string {
	return fmt.Sprintf(`services:
    mysql:
        image: mysql:8.0
        container_name: %s
        restart: unless-stopped
        environment:
            MYSQL_ROOT_PASSWORD: %s
            MYSQL_DATABASE: %s
        ports:
            - "%d:3306"
        volumes:
            - %s

    redis:
        image: redis:7.0
        container_name: %s
        restart: unless-stopped
        command: ["redis-server", "--requirepass", "%s"]
        ports:
            - "%d:6379"
        volumes:
            - %s

volumes:
    mysql_data:
    redis_data:
`, cfg.MySQL.ContainerName, cfg.MySQL.RootPassword, cfg.MySQL.Database, cfg.MySQL.Port, cfg.MySQL.Volume,
		cfg.Redis.ContainerName, cfg.Redis.Password, cfg.Redis.Port, cfg.Redis.Volume)
}
