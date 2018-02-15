
<h1 align="center">Get Youtube Thumbnail</h1>
<!-- <div align="center">
<a href="https://github.com/goreleaser/goreleaser/releases/tag/v0.45.1">
<img src="https://img.shields.io/github/release/qubyte/rubidium.svg?style=flat-square" alt="latest-release"></a>
<a href="https://semaphoreci.com/docs/how-to-get-build-badge.html">
<img src="https://img.shields.io/travis/rust-lang/rust.svg?style=flat-square" alt="build-passing"></a>
<a href="https://codecov.io/">
<img src="https://img.shields.io/scrutinizer/coverage/g/filp/whoops.svg?style=flat-square" alt="code-coverage"></a>
</div>
<hr> -->

By default, get_youtube_thumbnail gets a max resolution image, if image has been not found, get_youtube_thumbnail try to getting a high resolution image.


## How To Install

There are 2 ways to install "get_youtube_thumbnail". First way is simply to download ready build for your OS and the second way if you have installed golang on your system use "go get" then make build from source 

- Download ready binary for your OS:

     [Latest Release](#)
     <!-- https://github.com/goreleaser/goreleaser/releases/tag/v0.45.1 -->

 - or use go get command:
```
go get github.com\iqhater\get_youtube_thumbnail
```
then build your source
```
go build
```
 

## How to Use

1. Run binary file in console. (for linux and darwin os). For windows run `get_youtube_thumbnail.exe`
```
./get_youtube_thumbnail
```
2. Paste youtube url link
```
$ ./get_youtube_thumbnail

Enter Youtube Url: https://www.youtube.com/watch?v=glHtYwHidUY
```
3. Profit;) Your image saved in created "thumbnails" directory
```
$ ./get_youtube_thumbnail

Enter Youtube Url: https://www.youtube.com/watch?v=glHtYwHidUY

Already Done:)
```

## License

[MIT License](LICENSE)