const dropAreaList = Array.from(document.getElementsByClassName("drop-area"))

dropAreaList.forEach((dropArea) => {
    const fileInput = dropArea.querySelector('input[type="file"]')

    if (!fileInput) {
        return
    }

    fileInput.addEventListener("change", () => {
        dropArea.classList.add("uploaded")
    })

    dropArea.addEventListener("dragenter", (event) => {
        event.preventDefault()
        event.stopPropagation()
        dropArea.classList.add("hover")
    })
     
    dropArea.addEventListener("dragover", (event) => {
        event.preventDefault()
        event.stopPropagation()
        dropArea.classList.add("hover")
    })
     
    dropArea.addEventListener("dragleave", (event) => {
        event.preventDefault()
        event.stopPropagation()
        dropArea.classList.remove("hover")
    })
     
    dropArea.addEventListener("drop", (event) => {
        event.preventDefault()
        event.stopPropagation()
        dropArea.classList.remove("hover")
        dropArea.classList.add("uploaded")
     
        if (!event.dataTransfer.files.length) {
            return
        }

        fileInput.files = event.dataTransfer.files
    })
})


