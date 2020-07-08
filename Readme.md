### golang: comparison of libraries pgx, sqlx, gorm
  
   libraries:

``` 
   "github.com/jackc/pgx" 
```
![bench](picture/bench_pgx.png)

```
   "github.com/jmoiron/sqlx" with "github.com/jackc/pgx"
```
![bench](picture/bench_sqlx_pgx.png)

```
	"github.com/jmoiron/sqlx" with "github.com/lib/pq"
```
![bench](picture/bench_sqlx_pq.png)

```
   "github.com/jinzhu/gorm"
```
![bench](picture/bench_gorm.png)