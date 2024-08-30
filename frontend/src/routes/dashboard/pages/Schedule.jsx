import { PiUser } from "react-icons/pi";
import Input from "../../../components/input/Input"

const Schedule = () => {
  return (
    <div className="flex-1 flex items-center justify-center">
      <div className="w-3/4 h-4/5 rounded-md bg-gray-700 shadow-lg shadow-gray-400">
      
        <form className="border w-full h-full flex justify-center">

          <div className="w-2/3">

            <Input 
              title="Nome"
              placeholder="Escreva o nome do paciente"
              name="name"
              icon={<PiUser />}
              shadowLess={true}
            />

            <Input 
              title="Horário"
              placeholder="Escreva o nome do paciente"
              name="name"
              icon={<PiUser />}
              shadowLess={true}
            />
            
            <Input 
              title="Horário"
              placeholder="Escreva o nome do paciente"
              name="name"
              icon={<PiUser />}
              shadowLess={true}
            />

          </div>

        </form>

      </div>
    </div>
  )
}

export default Schedule