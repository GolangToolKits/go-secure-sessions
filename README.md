# go-secure-sessions
A new session module to replace gorilla/sessions.
This go-secure-sessions is not dependant on any router

Uses AES-128, AES-192, or AES-256 to store cookie sessions.


### Using go-secure-sessions

```go

    import "github.com/GolangToolKits/go-secure-sessions"

    // securekey must be at least 16 char long
    // The key argument should be the AES key,
    // either 16, 24, or 32 bytes to select
    // AES-128, AES-192, or AES-256.
    var secretKey = "dsdfs6dfs61dssdfsdfdsdsfsdsdllsd"

    var cf ConfigOptions
	cf.maxAge = 3600
    cf.path= "/"
	sessionManager, err := NewSessionManager(secretKey, cf)
	if err != nil {
		fmt.Println(err)
	}

    r, _ := http.NewRequest("POST", "/test/test1", nil)
    var w http.ResponseWriter

    // If a session cookie already exists in the request, it is loaded
    // Otherwise a completely new session is created with the name given
    session := sessionManager.NewSession(r, "new_test_sesstion")


    type SomeObject struct{
        Id int
        Name string
    }

    obj := SomeObject{
        Id: 1,
        Name: "test"
    }
    
    // needed to serialize and deserialize this object
    gob.Register(SomeObject{})

    //set some session values
    session.Set("test1", "some test1 value")
    session.Set("test2", "some test2 value")
    session.Set("test3", obj)
    // save the session before quitting or the values will be lost
    // session is saved securily as a cookie in the user's browser
    err:= session.Save(w)
    if err != nil{
        log.Println("Sesion not saved")
    }

    // Read a value out of the session
    v1:= session.Get("test1")


```