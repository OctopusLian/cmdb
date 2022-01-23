function request(method, url, params, callback) {
    jQuery.ajax({
        type: method,
        url : url,
        data: params,
        beforeSend: function(xhr) {
            xhr.setRequestHeader("X-Xsrftoken", jQuery.base64.atob(jQuery.cookie("_xsrf").split("|")[0]));
        },
        success: function(response) {
            switch(response["code"]) {
                case 200:
                    callback(response);
                    swal({
                        title: response["text"],
                        text: '',
                        type: "success",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    break;
                case 400:
                    var errors = [];
                    jQuery.each(response["result"], function(k, v) {
                        errors.push(v['Message']);
                    });
                    if(!errors) {
                        errors.push(response["text"]);
                    }
                    swal({
                        title: '',
                        text: errors.join("\n"),
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });

                    break;
                case 403:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    }, function() {
                        window.location.replace("/");
                    });
                    break;
                case 500:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    break;
                default:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    break;
            }
        },
        dataType: "json"
    });
}