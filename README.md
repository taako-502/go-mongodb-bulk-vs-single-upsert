# Go MongoDB Upsert vs UpsertMany

## Usage

### 通常実行

```zsh
make run
```

### ベンチマーク測定

```zsh
make benchmark
```

## 実行結果
### グラフ

![upsert_speed](https://github.com/taako-502/go-mongodb-upsert-vs-upsertmany/assets/36348377/5aee6efc-2165-4298-b258-2d06485fc697)

### 生データ

| Count   | Upsert[ms] | Upsert with BulkWrite[ms] |
|---------|------------|---------------------------|
| 2       | 1.622125   | 0.99875                   |
| 10      | 3.543333   | 1.282958                  |
| 100     | 21.497917  | 1.559125                  |
| 1000    | 235.175125 | 14.436208                 |
| 10000   | 2490.053541| 123.69225                 |
| 100000  | 27616.708  | 1228.453167               |
| 1000000 | 252393.2512| 11015.91683               |
