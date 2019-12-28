# voxelize

### ディレクトリの解説

```text
├── adapter 外部接続するインターフェイス部分
│
├── asset   データ等の資材
│
├── loader  外部ファイルを読み込むローダ
│   └── obj .obj形式を読み込むローダ
│
├── model   ボクセル化の過程で中心になるオブジェクト
    │       voxelization.goがボクセル化のファサード(呼び出し口)
    ├── cell.go
    ├── color.go
    ├── lattice.go
    ├── obje.go
    ├── point.go
    ├── voxel.go
    └── voxelization.go
```

        
