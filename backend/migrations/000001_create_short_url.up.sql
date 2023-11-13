CREATE TABLE IF NOT EXISTS URL (
    short_url PRIMARY KEY,  
    long_url NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE ,
    expires_at DATE NOT NULL DEFAULT (CURRENT_DATE + interval '90 days') ,
    last_visited DATE NOT NULL DEFAULT CURRENT_DATE ,

)