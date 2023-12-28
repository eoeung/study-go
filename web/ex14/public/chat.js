$(function(){
    // 현재 브라우저가 EventSource를 지원하는지 확인
    // https://developer.mozilla.org/en-US/docs/Web/API/EventSource#browser_compatibility
    if(!window.EventSource){
        alert("No EventSource");
        return;
    }

    const $chatlog = $('#chat-log');
    const $chatmsg = $('#chat-msg');

    // 빈 값 유효성 검사
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
            name: username,
            msg: $chatmsg.val(),
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
    // new EventSource(url, options)
    const es = new EventSource('/stream'); // 서버에서 만들어 놓은 EventSource 인스턴스와 연결
    // EventSource 연결되면
    es.onopen = function(e){
        // 유저 추가
        $.post("/users", {
            name: username
        })
    }

    // EventSource를 통한 데이터 수신
    es.onmessage = function(e){
        msg = JSON.parse(e.data);
        addMessage(msg);
    }

    // unload 전 이벤트
    // Ex) 브라우저 닫기
    window.onbeforeunload = function(){
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        });

        es.close();
    }
})