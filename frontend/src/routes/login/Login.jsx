import { PiUser } from "react-icons/pi";
import { GoLock } from "react-icons/go";
import logo from "../../assets/logo.svg"
import Input from "../../components/input/Input";

const Login = () => {
  return (
    <main className="flex h-screen w-full p-5 gap-6">

      <div className="w-2/5 bg-montain bg-center rounded-lg shadow-lg shadow-slate-300"></div>
      <div className="w-3/5 flex flex-col items-center justify-center gap-5">
        <div className="w-1/4">
          <img src={logo} alt="" />
          <p className="text-center text-sm">Seja bem-vindo(a)</p>
        </div>
          
        <form className="w-3/5 flex flex-col gap-5">
          <Input 
            icon={<PiUser />}
            name="login"
            placeholder="Digite aqui o login de acesso"
            title="Login"
          />

          <Input 
            icon={<GoLock />}
            name="login"
            placeholder="Digite aqui a senha de acesso"
            title="Senha"
            isPass={true}
          />

          <button 
            type="submit" 
            className="bg-zinc-700 text-white font-bold py-2 rounded-md mt-5"
          >
            Entrar
          </button>

        </form>
      </div>

    </main>
  )
}

export default Login