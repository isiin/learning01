## create vite 

```bash
pnpm create vite
  Project name:
│  react-ts
│
◇  Select a framework:
│  React
│
◇  Select a variant:
│  TypeScript + SWC
│
◇  Use Vite 8 beta (Experimental)?:
│  No
│
◇  Install with pnpm and start now?
│  Yes
```

## biome

### eslint削除

1. package.jsonから削除
- @eslint/js
- eslint
- eslint-plugin-react-hooks
- eslint-plugin-react-refresh
- globals
- typescript-eslint

2. 更新
```bash
pnpm install
```

### biome追加

```bash
pnpm add -D -E @biomejs/biome

pnpm exec biome init
```

### コマンド修正 

1. package.json修正
```json
  ...
  "scripts": {
    ...
    "check": "biome check .",
    "fix": "biome check --write .",
  },
  ...
```

2. お試し
```bash
pnpm check 
```

### VSCode

1. .vscode/extensions.json
```json
{
    "recommendations": [
        "biomejs.biome"
    ]
}
```
2. .vscode/settings.json
```json
{
  "editor.codeActionsOnSave": {
    "source.fixAll.biome": "explicit",
  },
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "biomejs.biome"
}
```

## MUI

1. パッケージ追加
```bash
# 基本
pnpm add @mui/material @emotion/react @emotion/styled
# Robotoフォント
pnpm add @fontsource/roboto
# Materialアイコン
pnpm add @mui/icons-material
```

2. main.tsxにRobotoフォントのimport追加
```tsx
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
```

## Leaflet

```bash
pnpm add leaflet leaflet.markercluster
pnpm add -D @types/leaflet @types/leaflet.markercluster
pnpm add -D @types/geojson
```

## React Router

```bash
pnpm add react-router
pnpm add -D @types/react-router
```

## 状態の分類と保存先
1. 一時的なUI状態
- 具体例: モーダルの開閉、入力中の文字、スクロール位置	
- 推奨される保存先: React State (useState)
- 理由: 画面をリロードしたらリセットされて良い（されるべき）ため。
2. ユーザー設定
- 具体例: 地図の種類（航空写真など）、ダークモード、表示件数
- 推奨される保存先: LocalStorage
- 理由	ブラウザを閉じても設定を覚えておいてほしいため。
3. 共有したい状態
- 具体例: 選択中の作業区、検索条件、ページ番号	
- 推奨される保存先:URLパラメータ (Router)	
- 理由: URLをコピーして同僚に送ったとき、同じ画面（同じ作業区が開いた状態）を再現できるため。リロードしても消えない。
4. 認証情報	
- 具体例: アクセストークン、セッションID	
- 推奨される保存先: Cookie (HttpOnly)	
- 理由: セキュリティ（XSS対策）のため。次点でLocalStorage。
5. サーバーデータ
- 具体例: 顧客リスト、マスタデータ	
- 推奨される保存先: メモリキャッシュ (React Query等)	
- 理由: サーバーにあるデータが「正」であり、フロントは一時的に持っているだけだから。

## Reactのステート管理ライブラリ（Zustand, Jotai, Reduxなど）

### 主な使用目的（導入する理由）
1. Propsのバケツリレー（Prop Drilling）の解消
- 課題: 親コンポーネントにあるデータを、深い階層の孫コンポーネントで使いたい場合、間のコンポーネント全てに props を渡していく必要があります。
- 解決: ステート管理ライブラリを使うと、「ストア（Store）」 というコンポーネントツリーの外側にある場所にデータを置けます。どのコンポーネントからでも、直接そのデータを取りに行けるため、親から子へ延々と渡す必要がなくなります。
2. 再レンダリングのパフォーマンス最適化
- 課題: Reactの useContext は便利ですが、Context内の値が1つでも変わると、そのContextを使っている全てのコンポーネントが再レンダリングされてしまいます。
- 解決: ZustandやReduxなどは、ストアの中の**「自分が必要なデータだけ」** を監視する仕組み（Selector）を持っています。関係ないデータが更新されても、自分のコンポーネントは再レンダリングされないため、パフォーマンスが向上します。
3. 状態ロジックの分離と整理
- 課題: コンポーネントの中に useState や useEffect が増えすぎると、UI（見た目）のコードとロジック（処理）のコードが混ざり合い、読みづらくなります。
- 解決: 「データをどう更新するか（アクション）」をストア側に記述することで、コンポーネント側は「データを表示する」「アクションを呼ぶ」ことだけに集中でき、コードがすっきりします。
4. 状態の永続化やデバッグ
- 課題: ブラウザをリロードしてもデータを保持したい（localStorageへの保存）場合や、バグ調査のために「いつ、どうやってデータが変わったか」を知りたい場合、自前で実装するのは大変です。
- 解決: 多くのライブラリには、設定一つで localStorage に同期する機能（ミドルウェア）や、ブラウザの拡張機能で状態の履歴を見られるツール（Redux DevToolsなど）が用意されています。

### 各ライブラリの簡単な特徴
1. Redux (Redux Toolkit):
- 最も有名で歴史がある。
- ルールが厳格でコード量は増えがちだが、データの流れが追いやすく、大規模チーム開発に向いている。
2. Zustand:
- 非常にシンプルで軽量。
- Hooksベースで書けるため、useState の延長のような感覚で簡単に導入できる。現在のReact開発で非常に人気。
3. Jotai / Recoil:
- Atom（アトム） と呼ばれる小さな単位で状態を管理する。
- 状態同士の依存関係（Aが変わったらBも自動で計算し直すなど）を定義するのが得意。
