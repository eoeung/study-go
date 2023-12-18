$(function(){
    if(!window.EventSource){
        alert("No EventSource");
        return;
    }

    const $chatlog = $('#chat-log');
    const $chatmsg = $('#chat-msg');

    const isBlank = function(string){
        return string == null || string.trim() === "";
    };

    let username;
    while (isBlank(username)){
        username = prompt("What's your name?");
        if(!isBlank(username)){
            $('#user-name').html('<b>' + username + '</b>');
        }
    }

    // 채팅 메시지 전송
    $('#input-form').on('submit', function(e){
        $.post('/messages', {
            msg: $chatmsg.val(),
            name: username
        });
        $chatmsg.val("");
        $chatmsg.focus();

        return false;
    });

    // 채팅 로그에 메시지 출력
    const addMessage = function(data){
        let text = "";
        if(!isBlank(data.name)){
            text = '<strong>' + data.name + ': </strong>';
        }

        text += data.msg;
        $chatlog.prepend('<div><span>' + text + '</span></div>');
    };

    // addMessage() 함수 동작 테스트
    // addMessage({
    //     msg: 'hello',
    //     name: 'aaa',
    // })

    // addMessage({
    //     msg: 'hello2',
    // })

    // 이벤트소스 생성
    const es = new EventSource('/stream')
    es.onopen = function(e){
        // 유저 추가
        $.post("/users", {
            name: username
        })
    }

    es.onmessage = function(e){
        msg = JSON.parse(e.data);
        addMessage(msg);
    }

    window.onbeforeunload = function(){
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        });

        es.close();
    }
})