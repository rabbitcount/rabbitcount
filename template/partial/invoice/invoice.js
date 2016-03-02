//function init() {
//    $("#content-wrapper").on("click", "button", function () {
//        var formData = new FormData($("#invoice-add")[0])
//
//        $.ajax({
//            url: '/api/mongo',
//            type: 'POST',
//            data: formData,
//            async: false,
//            cache: false,
//            contentType: false,
//            processData: false,
//            success: function (returndata) {
//                alert(returndata);
//            },
//            error: function (returndata) {
//                alert(returndata);
//            }
//        });
//    });
//}

function createInvoice() {
    alert("12");

    var formData = new FormData($("#invoice-add")[0])
    //$.post("/api/user/create", formData);
    $.ajax({
        url: '/api/mongo',
        type: 'POST',
        data: formData,
        async: false,
        cache: false,
        contentType: false,
        processData: false,
        success: function (returndata) {
            alert(returndata);
        },
        error: function (returndata) {
            alert(returndata);
        }
    });
}
