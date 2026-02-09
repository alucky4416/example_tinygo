# tinygo_keeb/zero-kb02 でSHT3x温湿度センサーからデータ取得して表示するサンプル
  
## 概要
　tinygo_keeb/workshopで組み立てる「zero-kb02」（自作キーボード／マクロパッド）には、GROVEコネクタがついており、ここにGROVE(I2C)対応のさまざまなセンサーをつなぐことができます。    

![zero-kb02 and SHT31](https://github.com/alucky4416/example_tinygo/example_i2c_sht3x_ws_rp2040/kb02_and_SHT31.jpg)

　tinygo_keeb/workshopで公開されている「zero-kb02」用のサンプルには、sht40温湿度センサー用のものがあり、「zero-kb02」の液晶に温度・湿度表示することまでてきてしまいます。
手元にSHT31があったので、学習と動作確認を兼ねて、SHT31用に修正、Printfだけの単純なものに修正してみました。
I2Cデータ取得とPrintfだけにしてしまったので、コントローラとSHT31センサーがあれば、とりあえず、動かすことができます。

https://github.com/tinygo-keeb/workshop

## ファームウェア(uf2)の焼き方
　このサンプルのビルド済みのファームウェア(sht31_wsrp2040zero.uf2)をソースと一緒に登録しています。
焼き方はtinygo_keeb/workshopに詳しいビルド方法の説明があります。(tinygo_keebの皆様、ありがとう！)

### ざっくり説明
　waveshare-rp2040-zeroのBOOTボタンを押しながら、PCのUSBポートに接続すると、RP2という名前のUSBメモリーとして認識されます。そこにuf2ファイルをコピーするだけで、waveshare-rp2040-zeroのファームウェアとして登録されます。  
　書き込んだ後、USBメモリーは自動的に接続解除され、その後すぐに(waveshare-rp2040-zeroは)再起動して書き込んだファームウェアが動き始めます。PCの仮想COMポート(232C)として認識されます。デバイスマネージャーなどで確認できます。  
　このサンプルでは、SHT31から取得した温度と湿度を、Printfで出力していますが、出力先は、この仮想COMポート(232C)となっています。
「tinygo monitor」オプションを使うと、仮想COMポートから受信したテキストをターミナルに表示させることができます。
  
　注意、uf2ファイルがちゃんと書かれたかどうか確認したくなる人もいると思いますが、再度、BOOTボタンを押しながらPCに接続してしまうと、前に書き込んでいたファームウェアは削除されてしまいますので、注意してください。
　一度、書き込んだファームウェアは、BOOTボタンを押しながら再接続しない限り、電源断、RESETボタン押す、USBポートからとりはずしても消えることはありません。

# おまけ
　ターミナルにテキスト表示させるだけでは寂しいので、LabVIEWでデータを表示するサンプルを作ってみました。
LabVIEW 2020以降、NI-VISAがインストールされたPCで動作します。(基本的にWindowsPC)

以下、サンプル画面

![LabVIEW Sample](https://github.com/alucky4416/example_tinygo/example_i2c_sht3x_ws_rp2040/example_tinygo_sht31_ws_rp2040zero_labview.PNG)

グラフに変化をつけるためにセンサーを手で抑えたり放したりしています。センサーは正常です。  
画面左上の「VISA resource name」はPCに存在するCOMポート一覧を選択できるリストボックスになっていますので、そこから、該当するCOMポートを選択してから、Runボタンで実行します。
