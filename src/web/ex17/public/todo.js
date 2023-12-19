(function($) {
    'use strict';
    $(function() {
      var todoListItem = $('.todo-list');
      var todoListInput = $('.todo-list-input');
      $('.todo-list-add-btn').on("click", function(event) {
        event.preventDefault();
  
        var item = $(this).prevAll('.todo-list-input').val();
        
        // // 추가할 아이템을, 프론트에 출력하기만 하는 것이 아닌 서버로 전송
        // if (item) {
        //   $.post("/todos", {name: item}, function(e){
        //     // 서버로부터 응답이 왔을 때, add
        //     addItem({name: item, completed: false});
        //   })
        //   // todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox'/>" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        //   todoListInput.val("");
        // }

        if (item) {
          $.post("/todos", {name: item}, addItem) // 서버에서 넘어온 값을 바로 사용
          // todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox'/>" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
          todoListInput.val("");
        }
      });

      const addItem = function(item){ // json object가 올 예정
        // console.log(item)
        if(item.completed){
          todoListItem.append("<li class='completed'" + "id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked=''/>" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        } else{
          todoListItem.append("<li id='" + item.id +  "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox'/>" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
        }
        todoListInput.val("");
      }

      // 모든 목록을 불러오기
      $.get("/todos", function(items){
        items.forEach(elem => {
          addItem(elem);
        });
      })
  
      // 서버에 요청을 보내서, completed면 화면에서 바뀌도록 수정
      todoListItem.on('change', '.checkbox', function() {
        const id = $(this).closest("li").attr('id');
        const $self = $(this);
        let complete = true;
        if ($self.attr('checked')){
          complete = false;
        }

        $.get("/complete-todo/" + id + "?complete=" + complete, function(data){
          if(complete){ // 미완료 → 완료 상태로 요청하는 경우
            $self.attr('checked', 'checked');
          } else{
            $self.removeAttr('checked');
          }
    
          $self.closest("li").toggleClass('completed');
        })
      });
  
      todoListItem.on('click', '.remove', function() {
        // url: todos/id
        // method: DELETE
        const id = $(this).closest("li").attr('id');
        const $self = $(this);

        $.ajax({
          url: "/todos/" + id,
          type: "DELETE",
          success: function(data){
            if(data.success){
              $self.parent().remove();
            // $(this).parent().remove(); // function의 this가 될 수 있음
            }
          }
        })
      });
  
    });
  })(jQuery);