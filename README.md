
# Graded Challenge 3 - P2

Benedict Kevin Sofyan - Instaclone Graded Challenge

## Demo 

the app deployed to heroku and can be accessible with this url
```
https://instaclone-gc-d3dd21c2b9a4.herokuapp.com/swagger/index.html
```

use this sample username for login
```
{
  "email": "kevinsofyan.13@gmail.com",
  "password": "password123"
}
```

## Database

this app use supabase postgre as database
use this env to run in local

```
DB_HOST=aws-0-ap-southeast-1.pooler.supabase.com
DB_USER=postgres.eayclkabtrbpmtybtynf
DB_PASSWORD=asdQWE123!
DB_NAME=postgres
DB_SCHEMA=instaclone                   
DB_PORT=5432    
```

## Running in local

to run this in local just use this command

```
go run main.go
```

the default port will be :8080