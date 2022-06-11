const data = {
    "root": [
        {"method": "GET", "path": "/", "parameters": {}, "body": {}, "authorization": ""}
    ],

    "api": [
        {"method": "GET", "path": "/api", "parameters": {}, "body": {}, "authorization": ""}
    ],

    "models": [
        {"method": "GET", "path": "/models/user",
            "parameters": {},
            "body": {},
            "authorization": ""
        },
        {"method": "GET", "path": "/models/user/{id}",
            "parameters": {"id":""},
            "body": {},
            "authorization": ""
        },
        {"method": "POST", "path": "/models/user/create",
            "parameters": {},
            "body": {"username":"", "password":"", "name":"", "email":""},
            "authorization": ""
        },
        {"method": "PUT", "path": "/models/user/update/{id}",
            "parameters": {"id":""},
            "body": {"username":"", "password":"", "name":"", "email":""},
            "authorization": ""
        },
        {"method": "DELETE", "path": "/models/user/delete/{id}",
            "parameters": {"id":""},
            "body": {},
            "authorization": ""
        }
    ],

    "auth": [
        {"method": "POST", "path": "/auth/sign-up",
            "parameters": {},
            "body": {"username":"", "password":"", "name":"", "email":""},
            "authorization": ""
        },
        {"method": "POST", "path": "/auth/sign-in",
            "parameters": {},
            "body": {"username":"", "password":""},
            "authorization": ""
        },
        {"method": "POST", "path": "/auth/sign-out",
            "parameters": {},
            "body": {},
            "authorization": ""
        },

        {"method": "POST", "path": "/auth/token",
            "parameters": {},
            "body": {},
            "authorization": ""
        },
        {"method": "POST", "path": "/auth/sessions",
            "parameters": {},
            "body": {"username":"", "password":"", "name":"", "email":""},
            "authorization": ""
        }
    ],

    "test": [
        {"method": "GET", "path": "/test/public",
            "parameters": {},
            "body": {},
            "authorization": ""
        },
        {"method": "POST", "path": "/test/private-free",
            "parameters": {},
            "body": {},
            "authorization": ""
        },
        {"method": "POST", "path": "/test/private-premium",
            "parameters": {},
            "body": {},
            "authorization": ""
        },
        {"method": "POST", "path": "/test/private-admin",
            "parameters": {},
            "body": {},
            "authorization": ""
        }
    ]
}