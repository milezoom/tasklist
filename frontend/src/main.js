let btnAdd = document.getElementById("btn-add")
let taskDesc = document.getElementById("task-desc")
let taskTime = document.getElementById("task-time")
let taskList = document.getElementById("task-list")

window.addTask = function () {
  let desc = taskDesc.value
  let time = taskTime.value

  let result = window.storeTask(desc, time)
  console.log(result)
  if (result == "") {
    window.prependTask(desc, time)
    taskDesc.value = ""
    taskTime.value = ""
    taskDesc.focus()
  }
}

window.prependTask = function (desc, time) {
  let taskDom = document.createElement("div")
  taskDom.setAttribute("class", "columns")

  let taskDescDom = document.createElement("p")
  taskDescDom.innerText = desc
  taskDescDom.setAttribute("class", "column is-9")

  let taskTimeDom = document.createElement("p")
  taskTimeDom.innerText = time
  taskTimeDom.setAttribute("class", "column is-3")

  taskDom.append(taskDescDom, taskTimeDom)
  taskList.prepend(taskDom)
}

window.storeTask = function (desc, time) {
  let result = ""
  let str = JSON.stringify({
    description: desc,
    time: time
  })

  window.go.main.App.StoreTask(str).then((resp) => {
    result = resp
  })

  return result
}

window.loadTask = function () {
  window.go.main.App.GetTasks().then((resp) => {
    console.log(resp)
    if (resp != "false") {
      let tasks = JSON.parse(resp)
      taskList.innerHTML = ""
      tasks.forEach(element => {
        window.prependTask(element.description, element.time)
      })
    }
  })
}

window.addEventListener("load", function (e) {
  window.loadTask()
})

taskTime.addEventListener("keypress", function(e) {
  if (e.key == 'Enter') {
    window.addTask()
  }
})