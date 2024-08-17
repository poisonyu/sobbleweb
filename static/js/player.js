

window.onload = function() {
	const dp = new DPlayer({
        container: document.getElementById('dplayer'),
        autoplay: false,
        theme: '#FADFA3',
        loop: true,
        lang: 'zh-cn',
        screenshot: true,
        hotkey: true,
        //chromecast: true,
        preload: 'auto',
        //logo: 'logo.png',
        volume: 0.7,
        mutex: true,
        video: {
            url: '/static/demo.mp4',
            //pic: 'dplayer.png',
            //thumbnails: 'thumbnails.jpg',
            type: 'auto',
        },
        contextmenu: [
            {
                text: 'custom1',
                link: 'https://github.com/DIYgod/DPlayer',
            },
            {
                text: 'custom2',
                click: (player) => {
                    console.log(player);
                },
            },
        ],
        highlight: [
            {
                time: 20,
                text: '这是第 20 秒',
            },
            {
                time: 120,
                text: '这是 2 分钟',
            },
        ],
    });
}

function startplay() {
	let url = document.querySelector("#kw").value;
	if (url != null) {
		dp.switchVideo(
			{
				url: url,
			}
		)
		dp.play();
	} else {
		console.log('url is null', url);
		alert('url is null')
	}
}

function clear() {
	document.querySelector('#kw').value = '';
}

