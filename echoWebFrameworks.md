# Echo Framework

Echo is a high-performance, minimalist Go web framework designed for building scalable and maintainable server-side applications.

---

## REST (Representational State Transfer)

REST is an architectural style used to design web services.  
It leverages HTTP methods to perform CRUD (Create, Read, Update, Delete) operations on resources.

### Key Principles of REST
- **Stateless** – Every request is independent.
- **Client-Server Architecture** – Separation of concerns.
- **Uniform Interface** – Standardized communication using HTTP.
- **Cacheable** – Response data can be cached.
- **Layered System** – Intermediaries like proxies, load balancers allowed.
- **Resource-based** – Everything treated as a resource accessed via URLs.

# Content

---

## Terms

- `Serialization` (`struct → JSON/string/binary`)
- `Marshalling` = converting a Go struct → JSON
- `Deserialization` (`JSON/string/binary → struct/object`)
- `Unmarshalling` = converting JSON → Go struct

## Methods

- **GET**  
  Ye request sirf data **fetch** ya **read** karne ke liye use hoti hai.  
  Server ki state change nahi hoti.

- **POST**  
  Ye request **naya data create** karne ke liye use hoti hai.  
  Har POST request se generally naya resource ban jata hai.

- **PUT**  
  Ye **pure resource ko update/replace** karne ke liye use hota hai.  
  Agar aapko **sirf ek field** update karni ho, tab bhi **pura payload (saare fields)** bhejne padte hain.

- **PATCH**  
  Ye **partial update** ke liye hota hai.  
  Yani **kisi single field** ko update karna ho to **sirf wahi field** bhejna hota hai,  
  saare fields dene ki jarurat nahi hoti.

- **DELETE**  
  Ye request kisi resource ko **delete/remove** karne ke liye use hoti hai.
