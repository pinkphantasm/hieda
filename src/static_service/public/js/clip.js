function callback() {
    Array.from(document.getElementsByClassName('clip-area')).forEach(function(clipArea) {
        const clipAreaText = clipArea.querySelector('.clip-area-text')
        const clipAreaButton = clipArea.querySelector('.clip-area-button')

        if (!clipAreaText || !clipAreaButton) {
            return
        }

        clipAreaButton.addEventListener('click', function() {
            navigator.clipboard.writeText(clipAreaText.innerText)
                .then(function() {
                    clipAreaButton.classList.add('success')
                    setTimeout(() => {
                        clipAreaButton.classList.remove('success')
                    }, 1500)
                })
                .catch(function() {
                    clipAreaButton.classList.add('error')
                    setTimeout(() => {
                        clipAreaButton.classList.remove('error')
                    }, 1500)
                })
        })
    })
}

const observer = new MutationObserver(callback)
observer.observe(document.getElementById('clip-area-wrapper'), {
    childList: true,
})
