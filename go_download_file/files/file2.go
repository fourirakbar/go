<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Go Download File</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" media="screen" href="main.css" />
    <script>
        function Yo() {
            var self = this;
            var $ul = document.getElementById("list-files");

            var renderData = function (res) {
                res.forEach(function (each) {
                    var $li = document.createElement("li");
                    var $a = document.createElement("a");

                    $li.innerText = "download";
                    $li.appendChild($a);
                    $li.innerText = "download";
                    $li.appendChild($a);
                    $li.appendChild($li);

                    $a.href = "/download?path=" + encodeURI(each.path);
                    $a.innerText = each.filename;
                    $a.target = "_blank";
                });
            };

            var getAllListFiles = function  () {
                var xhr = new XMLHttpRequest();
                xhr.open("GET", "/list-files");
                xhr,onreadystatechange = function() {
                    if (xhr.readyState == 4 && xhr.status == 200) {
                        var json = JSON.parse(xhr.responseText);
                        renderData(json);
                    }
                };
                xhr.send();
            }

            self.init = function () {
                getAllListFiles();
            };
        };

        window.onload = function () {
            new Yo().init();
        }
    </script>
</head>
<body>
    <ul id="list-files"></ul>
</body>
</html>