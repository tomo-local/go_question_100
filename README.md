# Go言語 100本ノック

Go言語の基礎から応用までを体系的に学習するための問題集です。
各問題には、問題ファイル (`questions/`) と解答例ファイル (`answers/`) が用意されています。

---

## 問題リスト

### Part 1: 基本 (1-20)
- [x] 1. `fmt.Println` vs `fmt.Printf`
- [x] 2. `len()` とマルチバイト文字
- [x] 3. `strings.Split`
- [x] 4. スライスと `append`
- [x] 5. `math.Ceil` vs `math.Floor`
- [x] 6. `time.Now` とフォーマット
- [x] 7. コマンドライン引数 (`os.Args`)
- [x] 8. 文字列と数値の変換 (`strconv`)
- [x] 9. スライスのソート (`sort.Ints`)
- [x] 10. 文字列の結合 (`strings.Join`)
- [x] 11. 標準入力の読み取り (`bufio.Scanner`)
- [x] 12. ファイルの読み書き (`os.ReadFile`, `os.WriteFile`)
- [x] 13. マップの基本 (作成、追加、アクセス)
- [x] 14. 構造体 (struct) の定義と利用
- [x] 15. Goの構造体をJSONに変換 (`json.Marshal`)
- [x] 16. JSONをGoの構造体に変換 (`json.Unmarshal`)
- [x] 17. HTTP GETリクエスト (`net/http.Get`)
- [x] 18. 乱数の生成 (`math/rand`)
- [x] 19. 文字列の検索 (`strings.Contains`, `strings.HasPrefix`)
- [x] 20. プログラムの一時停止 (`time.Sleep`)

### Part 2: 制御構文 (21-30)
- [x] 21. `if` と `else`
- [x] 22. `if` と短絡評価
- [x] 23. `switch` (基本)
- [x] 24. `switch` (式なし)
- [x] 25. `for` (3つのコンポーネント)
- [x] 26. `for` (while文のように)
- [x] 27. `for` (無限ループ)
- [x] 28. `for-range` (スライス)
- [x] 29. `for-range` (マップ)
- [x] 30. `defer` の基本

### Part 3: 関数とポインタ (31-40)
- [x] 31. 簡単な関数
- [x] 32. 複数の戻り値
- [x] 33. 名前付き戻り値
- [x] 34. 可変長引数
- [x] 35. 関数を値として渡す
- [x] 36. クロージャ
- [x] 37. ポインタの基本
- [x] 38. 関数にポインタを渡す
- [x] 39. `new()` 関数
- [x] 40. ポインタレシーバ vs 値レシーバ

### Part 4: メソッドとインターフェース (41-50)
- [x] 41. 構造体のメソッド
- [x] 42. 型にメソッドを追加
- [x] 43. インターフェースの基本 (`Stringer`)
- [x] 44. インターフェースを満たす
- [x] 45. 空のインターフェース (`interface{}`)
- [x] 46. 型アサーション
- [x] 47. 型スイッチ
- [x] 48. `error` インターフェース
- [x] 49. カスタムエラー
- [x] 50. `io.Reader` インターフェース

### Part 5: 並行処理 (51-65)
- [x] 51. goroutineの基本
- [x] 52. `sync.WaitGroup`
- [x] 53. チャネルの基本
- [x] 54. チャネルを使ったgoroutine間の通信
- [x] 55. チャネルの方向
- [x] 56. `range` と `close`
- [x] 57. `select` の基本
- [x] 58. `select` と `default`
- [x] 59. `time.Ticker`
- [x] 60. `time.Timer`
- [x] 61. バッファ付きチャネル
- [x] 62. `sync.Mutex`
- [x] 63. `sync.RWMutex`
- [x] 64. `sync.Once`
- [x] 65. `atomic` パッケージ

### Part 6: 標準ライブラリ (応用) (66-85)
- [ ] 66. `context` の基本
- [ ] 67. `context` とタイムアウト
- [ ] 68. `context` とキャンセル
- [ ] 69. `regexp` によるマッチ
- [ ] 70. `regexp` による置換
- [ ] 71. `flag` によるコマンドライン引数の解析
- [ ] 72. `net/http` でWebサーバを立てる
- [ ] 73. `net/http` の`HandleFunc`
- [ ] 74. `io.Copy`
- [ ] 75. `log` パッケージ
- [ ] 76. カスタムログフォーマット
- [ ] 77. `os` パッケージ (環境変数)
- [ ] 78. `os` パッケージ (ファイル操作)
- [ ] 79. `path/filepath` パッケージ
- [ ] 80. `encoding/csv` (書き込み)
- [ ] 81. `encoding/csv` (読み込み)
- [ ] 82. `time` のパースとフォーマット
- [ ] 83. `time` の加算と減算
- [ ] 84. `exec` による外部コマンドの実行
- [ ] 85. `url` パッケージ

### Part 7: テスト (86-95)
- [ ] 86. `testing` パッケージの基本
- [ ] 87. `t.Run` によるサブテスト
- [ ] 88. テーブル駆動テスト
- [ ] 89. ヘルパー関数
- [ ] 90. `TestMain`
- [ ] 91. ベンチマークテスト
- [ ] 92. `Example` 関数
- [ ] 93. アサーションライブラリを使わないテスト
- [ ] 94. モックの基本
- [ ] 95. インターフェースを使ったテスト

### Part 8: その他 (96-100)
- [ ] 96. タグ付き構造体
- [ ] 97. `init` 関数
- [ ] 98. ビルドタグ
- [ ] 99. `go generate`
- [ ] 100. `cgo` の簡単な例
