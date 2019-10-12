



## 针对sqlite:



```
if exist pokemon.db (
   del pokemon.db
)
if exist pokemon.db-shm (
   del pokemon.db-shm 
)
if exist pokemon.db-wal (
   del pokemon.db-wal  
)
sqlite3 pokemon.db ".read db.sql" 
pause
```