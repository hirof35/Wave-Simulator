# Pure Go Wave Simulator (Light & Sound)

Go言語の2Dゲームエンジン「Ebitengine」を使用し、**C言語コンパイラ（cgo）に一切依存しない（100% Pure Go）**で開発された、リアルタイムな光と音の波形シミュレーターです。

環境依存がないため、Windows、Mac、Linuxなどのあらゆる環境でコマンド一つで爆速ビルド・クロスコンパイルが可能です。
<img width="750" height="528" alt="スクリーンショット 2026-05-25 052633" src="https://github.com/user-attachments/assets/2d6fe9da-b297-4188-ba93-f5beb8018cdf" />

## 💡 特徴

- **cgo完全非依存 (Pure Go):** Windows環境などで発生しがちな GCC / MinGW のコンパイルエラーを完全に回避。
- **光と音の数理モデル:** - **音（周波数）:** 2つの異なるサイン波を合成し、音波の「うなり（Beats）」を表現。
  - **光（波長・スペクトル）:** 時間と位置の経過に応じてRGB（光の三原色）の値をサイン波で動的に変化させ、プリズムのような美しい虹色のグラデーションをリアルタイム生成。
- **高速なピクセル描画:** メモリ上のピクセルバッファ（`img.Pix`）を直接書き換えて画面に転送するため、非常に軽量かつ高速に動作します。

## 🛠️ 動作環境

- Go 1.16 以上
- OS: Windows, macOS, Linux (主要OSすべてに対応)

## 🚀 使い方

### 1. 依存関係のインストール
Ebitengine ライブラリを取得します。

```bash
go get [github.com/hajimehoshi/ebiten/v2](https://github.com/hajimehoshi/ebiten/v2)
2. アプリケーションの実行Bashgo run main.go
3. スタンドアロンアプリとしてビルド配布用の単一バイナリ（実行ファイル）を作成します。Windows用 (黒いコンソール画面を非表示にする設定):Bashgo build -ldflags="-H=windowsgui -s -w" -o wave_simulator.exe
Mac用:Bashgo build -ldflags="-s -w" -o wave_simulator_mac
📝 数理モデル・仕組み
1. 音波の合成（うなり）シミュレーター内の波の高さ（Y座標）は、以下の数式をベースに時間 $t$ と位置 $x$ で計算されています。
$$y = \sin(x \cdot \text{freq}_1 + t) + \sin(x \cdot \text{freq}_2 - t \cdot 0.5)$$これにより、2つの波の干渉による複雑な周期的なうねりが生まれます。
2. 光の色の表現波の各点のカラーは、光の波長変化をシミュレートするためにRGBそれぞれに異なる位相のサイン波を与えています。
Red: $127 + 127 \cdot \sin(x \cdot 0.01 + t)$Green: $127 + 127 \cdot \sin(x \cdot 0.015 + t)$Blue: $127 + 127 \cdot \cos(x \cdot 0.02 + t)$
📄 ライセンスThis project is licensed under the MIT License.
