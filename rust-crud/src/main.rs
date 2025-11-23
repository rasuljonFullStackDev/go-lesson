use actix_web::{web, App, HttpServer, Responder, HttpResponse};
use serde::{Deserialize, Serialize};
use rusqlite::{Connection, params};
use std::sync::Mutex;
use uuid::Uuid;

#[derive(Serialize, Deserialize)]
struct User {
    id: String,
    name: String,
    email: String,
}

struct AppState {
    db: Mutex<Connection>,
}

fn init_db() -> Connection {
    let conn = Connection::open("users.db").expect("Bazani ochib bo'lmadi");
    conn.execute(
        "CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT NOT NULL
        )",
        [],
    ).expect("Jadval yaratib bo'lmadi");
    conn
}

async fn get_users(data: web::Data<AppState>) -> impl Responder {
    let conn = data.db.lock().unwrap();
    let mut stmt = conn.prepare("SELECT id, name, email FROM users").unwrap();
    let users_iter = stmt.query_map([], |row| {
        Ok(User {
            id: row.get(0)?,
            name: row.get(1)?,
            email: row.get(2)?,
        })
    }).unwrap();

    let mut users = Vec::new();
    for user in users_iter {
        users.push(user.unwrap());
    }

    HttpResponse::Ok().json(users)
}

async fn get_user(data: web::Data<AppState>, path: web::Path<String>) -> impl Responder {
    let id = path.into_inner();
    let conn = data.db.lock().unwrap();
    let mut stmt = conn.prepare("SELECT id, name, email FROM users WHERE id = ?1").unwrap();

    let user = stmt.query_row(params![id], |row| {
        Ok(User {
            id: row.get(0)?,
            name: row.get(1)?,
            email: row.get(2)?,
        })
    });

    match user {
        Ok(u) => HttpResponse::Ok().json(u),
        Err(_) => HttpResponse::NotFound().body("Foydalanuvchi topilmadi"),
    }
}

async fn create_user(data: web::Data<AppState>, new_user: web::Json<User>) -> impl Responder {
    let conn = data.db.lock().unwrap();
    let id = Uuid::new_v4().to_string();
    conn.execute(
        "INSERT INTO users (id, name, email) VALUES (?1, ?2, ?3)",
        params![id, new_user.name, new_user.email],
    ).unwrap();

    HttpResponse::Created().json(format!("Yaratildi ID: {}", id))
}

async fn update_user(data: web::Data<AppState>, path: web::Path<String>, updated_user: web::Json<User>) -> impl Responder {
    let id = path.into_inner();
    let conn = data.db.lock().unwrap();

    let rows = conn.execute(
        "UPDATE users SET name = ?1, email = ?2 WHERE id = ?3",
        params![updated_user.name, updated_user.email, id],
    ).unwrap();

    if rows == 0 {
        HttpResponse::NotFound().body("Foydalanuvchi topilmadi")
    } else {
        HttpResponse::Ok().body("Yangilandi")
    }
}

async fn delete_user(data: web::Data<AppState>, path: web::Path<String>) -> impl Responder {
    let id = path.into_inner();
    let conn = data.db.lock().unwrap();
    let rows = conn.execute("DELETE FROM users WHERE id = ?1", params![id]).unwrap();

    if rows == 0 {
        HttpResponse::NotFound().body("Foydalanuvchi topilmadi")
    } else {
        HttpResponse::Ok().body("O'chirildi")
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let conn = init_db();
    let data = web::Data::new(AppState { db: Mutex::new(conn) });

    println!("🚀 Server ishga tushdi: http://127.0.0.1:8080");

    HttpServer::new(move || {
        App::new()
            .app_data(data.clone())
            .route("/users", web::get().to(get_users))
            .route("/users/{id}", web::get().to(get_user))
            .route("/users", web::post().to(create_user))
            .route("/users/{id}", web::put().to(update_user))
            .route("/users/{id}", web::delete().to(delete_user))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
