# HttpMessageServer
HttpMessageServer

### 실행
HttpMessageServer.exe

HttpMessageServer.exe -port 8080 -remote

HttpMessageServer.exe -task abc,efgh,dd

Linux

* HttpMessageServer.tar.gz 사용


### API

-사용 Method : GET, POST <br/>
-응답 Status Code : 200, 404 <br/>
-<a href="/status">http://localhost/status</a><br/>
-http://localhost/SetMessage?topic=aaa&message=12345678&task=processname</br>
-http://localhost/GetMessage?topic=aaa&timeout=1000</br>
-http://localhost/GetMessage?topic=aaa</br>
-http://localhost/GetLog?task=proc</br>
-http://localhost/ResetTopic?topic=aaa</br>


### Client
테스트는 브라우져에 위 URI를 입력 하여 확인 가능.

다음 항목들 API 를 찾아서 대충 코딩하여 사용하면 대충 굴러감...


c# 참고 

-HttpClient, FormUrlEncodedContent, KeyValuePair, PostAsync, GetAsync, Content.ReadAsStringAsync()

Golang 참고

-http.PostForm, ioutil.ReadAll, url.Values


-----------------------------------------------------------------------------------------------------------------------

결국, 문제는 죽느냐 사느냐 이것 하나 뿐이야.
