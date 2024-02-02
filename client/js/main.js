var menu, choose, dop, blockCat, elC,
menuVis = false, cats = [], logged = false, chooseVis = false, servLink = "http://localhost:3000";

function init(header) {
	console.log("ini");
	var main, filter = {};
	main = document.getElementsByClassName("main")[0];
	if(header == "Каталог") {
		filter = document.getElementsByClassName("filter")[0];
	}
	logged = localStorage.getItem("sec") ? 1 : 0;
	if(logged) {
		dop = `<div class="but buttonable" onclick="location.href='Profile.html';">Профиль</div>`;
		if(header == "Авторизация" || header == "Регистрация") {
			//location.href='Profile.html';
		}
	} else {
		dop = `<div class="but buttonable" onclick="location.href='Auth.html';">Вход</div>
		<div class="but buttonable" onclick="location.href='Reg.html';">Регистрация</div>`;
		if(header == "Профиль") {
			//location.href='Auth.html';
		}
	}
	main.innerHTML = mainHTML(header, main.innerHTML, filter.innerHTML);
	filter.innerHTML = "";
	if(header == "Каталог") initCatalog();
	if(header != "Главная") {
		menu = document.getElementsByClassName("burgerMenu")[0];
		menu.onmouseleave = function () {changeMenuVis();};
	} else {
		var menuI = document.getElementsByClassName("menu")[0];
		console.log(menuI.innerHTML);
		menuI.innerHTML += dop;
	}
	if(header == "Профиль") initProfile();
}

function initProfile() {
	var login = document.getElementById("login");
	sendToServerApp({}, 'GET', "profile/")
		.then(data => {
			console.log(data);
			if(data.status == 200){
				login.innerHTML = data.logname;
			} else {
				console.log("Нeуспешная загрузка профиля");
			}
		});
}

function initCatalog() {
	choose = document.getElementsByClassName("choose")[0];
	blockCat = document.getElementsByClassName("blockCat")[0];
	elC = document.getElementsByClassName("elC")[0];
	choose.onmouseleave = function () {changeChooseVis();};
	// addProduct("Пломбир Из Лавки cолёная карамель", 62, 75, "caram");
	// addProduct("Фруктовый лёд Вкусовые сосочки", 115, 75, "lyod");
	// addProduct("Пломбир с клубникой в глазури Кактус", 115, 80, "kaktyz");
	// addProduct("Смузи клубника", 1100, 1000, "smyzi");
	// addProduct("Джелато сливочное «Из Лавки» банан", 111, 90, "dzhelato");
	// addProduct("Мороженое фисташковое с миндалем", 249, 450, "mindal");
	// addProduct("Пломбир Банан трюфель 12% пакет", 249, 500, "banana");
	// addProduct("Чудеса света мороженое йогуртно-вишневое", 329, 450, "chudo");
	// addProduct("Эскимо Ассорти «Монстры»", 549, 620, "monster");
	// addProduct("Мороженое «Кокосовый мусс с шоколадом»", 65, 75, "muss");
	sendToServerApp({}, 'GET', "product/")
		.then(data => {
			console.log(data);
			if(data.status == 200){
				for(let product of data.products) {
					addProduct(product.zag, product.price, product.gram, product.img);
				}
			} else {
				console.log("Нeуспешная загрузка каталога");
			}
		});
}

function changeMenuVis() {
	menuVis = !menuVis;
	menu.dataset.vis = +menuVis;
}

function changeChooseVis(test) {
	chooseVis = !chooseVis;
	choose.dataset.vis = +chooseVis;
}

function changeZag(ev) {
	var answ = ev.target.parentElement.nextElementSibling;
	answ.dataset.vis = !(answ.dataset.vis === 'true');
}

function reg() {
	var inpL = document.getElementById("inpL"),
		inpP = document.getElementById("inpP");
	sendToServerApp({
        logname: inpL.value,
        password: inpP.value
    }, 'POST', "auth/register")
        .then(data => {
            console.log(data);
            if(data.status == 201){
				location.href='Auth.html';
            } else {
                console.log("Нeуспешная регистрация");
            }
        });
}

function vxod() {
	var inpL = document.getElementById("inpL"),
		inpP = document.getElementById("inpP");
	sendToServerApp({
        logname: inpL.value,
        password: inpP.value
    }, 'POST', "auth/login")
        .then(data => {
            console.log(data);
            if(data.status == 200){
                if(data.token) localStorage.setItem("sec", data.token);
				location.href='index.html';
            } else {
                console.log("Неверный логин или пароль ");
            }
        });
}

function sendToServerApp(bod, typeC, url) {
    let sed = {method: typeC};
    if (bod) {
        sed.headers = {'Content-Type': 'application/json'};
        if(localStorage.getItem("sec")) {
            sed.headers.Authorization = localStorage.getItem("sec");
            console.log("yyes send", localStorage.getItem("sec"));
        }
        sed.body = JSON.stringify(bod);
    }
    if(!url) url = "";
    return fetch(servLink + "/" + url, sed)
        .then(res => {
            if(res.status == 401 && localStorage.getItem("sec")) {
                localStorage.removeItem("sec");
            }
            if (!res.ok && res.status != 409) {
                throw new Error(`This is an HTTP error: The status is ${res.status}`);
            }
            return res.json();
        })
        .catch(data => data);
}

function addProduct(zag, price, gram, img) {
	var el = document.createElement("div");
	blockCat.appendChild(el);
	el.outerHTML = `
		<div class="elem buttonable">
			<div class="bgEl ${img}"></div>
			<div class="price">${price} ₽</div>
			<div class="textEl">
				${zag}
			</div>
			<div class="dopEl">${gram} гр</div>
		</div>
	`;
	cats.push(el);
	elC.innerText = cats.length;
}

function mainHTML(zag, main, filter) {
	var bool = zag == "Главная", boolCat = zag == "Каталог";
	return `
		<link rel="stylesheet" href="css/main.css">
		<div class="fullscreen-bg">
			<div class="over"></div>
		</div>
		<div class="columns${bool ? ' gl' : ''}">
			<div class="log">
				<div class="logo"></div>
				${bool ? '' : `<div class='logt'>${zag}</div>`}
				${boolCat ? filter : ''}
			</div>
			<div class="content${bool ? ' gl' : ' notGl'}${boolCat ? ' cat' : ''}">
				${main}
			</div>
			${bool ? '' : `
				<div class="burger buttonable" onclick="changeMenuVis();"></div>
			`}
		</div>
		<div class="footer">
			<div class="contact">
				© 1990–2023 «Ледяное волшебство»
				<div>magic@mail.com</div>
			</div>
			<div class="social">
				<div class="soc inst buttonable" onclick="location.href='https://instagram.com';"></div>
				<div class="soc yout buttonable" onclick="location.href='https://youtube.com';"></div>
				<div class="soc teleg buttonable" onclick="location.href='https://telegram.org/desktop';"></div>
			</div>
		</div>
		${bool ? '' : `
			<div class="burgerMenu">
				<div class="menuBox">
					<div class="but buttonable" onclick="location.href='index.html';">Главная</div>
					<div class="but buttonable" onclick="location.href='Catalog.html';">Каталог</div>
					<div class="but buttonable" onclick="location.href='FAQ.html';">FAQ</div>
					${dop}
				</div>
			</div>
		`}
	`;
}