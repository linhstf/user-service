package dataAcess

type DbContext struct {
    ConnectionString string
}

func Instant(connectionString string) *DbContext  {
    return &DbContext{
        ConnectionString: connectionString
    }
}

func Open()  {
    
}