
<h1 align="center">Get Youtube Thumbnail</h1>
<div align="center">
<a href="https://github.com/iqhater/get-youtube-thumbnail/releases/tag/v1.0.1">
<img src="https://img.shields.io/badge/Latest%20Release-v1.0.1-73C0E7.svg?style=flat-square" alt="latest-release"></a>
<a href='https://semaphoreci.com/iqhater/get-youtube-thumbnail'> <img src='https://semaphoreci.com/api/v1/iqhater/get-youtube-thumbnail/branches/master/badge.svg' alt='Build Status'></a>
<!-- <a href="https://codecov.io/">
<img src="https://img.shields.io/scrutinizer/coverage/g/filp/whoops.svg?style=flat-square" alt="code-coverage"></a> -->
</div>
<hr>

By default, get-youtube-thumbnail gets a max resolution image, if image has been not found, get-youtube-thumbnail try to getting a high resolution image.


## How To Install

There are 2 ways to install "get-youtube-thumbnail". First way is simply to download ready build for your OS and the second way if you have installed golang on your system use "go get" then make build from source 

- Download ready binary for your OS:

     [Latest Release](https://github.com/iqhater/get-youtube-thumbnail/releases/tag/v1.0.0)
     <!-- https://github.com/goreleaser/goreleaser/releases/tag/v0.45.1 -->

 - or use go get command:
```
go get github.com\iqhater\get-youtube-thumbnail
```
then build your source
```
go build
```
 

## How to Use

1. Run binary file in console. (for linux and darwin os). For windows run `get-youtube-thumbnail.exe`
```
./get-youtube-thumbnail
```
2. Paste youtube url link
```
$ ./get-youtube-thumbnail

Enter Youtube Url: https://www.youtube.com/watch?v=glHtYwHidUY
```
3. Profit;) Your image saved in created "thumbnails" directory
```
$ ./get-youtube-thumbnail

Enter Youtube Url: https://www.youtube.com/watch?v=glHtYwHidUY

Already Done:)
```

## License

[MIT License](LICENSE)