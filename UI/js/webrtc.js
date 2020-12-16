let constraints = {
    audio: true,
    video: [{
        frameRate: 30,
        width: {min: 1024, ideal: 1280, max: 1920},
        height: {min: 778, ideal: 720, max: 1080}
    }]
};

if (!navigator.getUserMedia) {
    alert('浏览器不支持WebRTC！');
} else {
    navigator.getUserMedia(constraints, onSuccess, onError);
}

function onSuccess(stream) {
    let video = document.querySelector('video');

    video.srcObject = stream;

    video.play();
}

function onError(error) {
    console.log('navigator.getUserMedia error:', error);
    alert(error);
}
