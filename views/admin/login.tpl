{{ template "templates/basic-layout.tpl" .}}

{{ define "content" }}
<div class="login">
    <form action="login" method="post">
        <input type="text" name="login" placeholder="Логин"/>
        <input type="password" name="password" placeholder="Пароль"/>
        <input type="submit" value="Войти">
    </form>
    <div>
        <a href="#">Забыли пароль?</a>
        <a href="#">Регистрация</a>
    </div>
</div>
{{end}}