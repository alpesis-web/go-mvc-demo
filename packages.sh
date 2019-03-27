packages=(
 "github.com/gorilla/mux" 
 "github.com/gorilla/sessions" 
 "github.com/go-redis/redis"
 "golang.org/x/crypto/bcrypt"
)


for i in "${packages[@]}"
  do
    go get "$i"; 
  done
