import { MdAutoGraph, MdEditDocument } from "react-icons/md";
import { IoIosBookmark } from "react-icons/io";
import { FaDollarSign } from "react-icons/fa";
import { FaUserGroup } from "react-icons/fa6";
import logo from "../../assets/logowhite.svg"
import Button from "../../components/button/Button";
import Board from "./pages/Board"
import Schedule from "./pages/Schedule"
import Register from "./pages/Register"
import Pacients from "./pages/Pacients"
import Payment from "./pages/Payment"
import { useRef, useState } from "react";

const Dashboard = () => {

  const btnList = [
    {
      icon: <MdAutoGraph />,
      title: "Painel",
      active: true,
      page: <Board />
    },
    {
      icon: <IoIosBookmark />,
      title: "Agendamento",
      page: <Schedule />
    },
    {
      icon: <MdEditDocument />,
      title: "Registrar",
      page: <Register />
    },
    {
      icon: <FaUserGroup />,
      title: "Pacientes",
      page: <Pacients />
    },
    {
      icon: <FaDollarSign />,
      title: "Pagamento",
      page: <Payment />
    },
  ]

  const [actualPage, setActualPage] = useState(null)

  const btnSetStatusList = useRef([])

  const handleAddSet = (s) => {
    btnSetStatusList.current.push(s)
  }

  const handleChangePage = p => {
    setActualPage(p)
  }

  return (
    <main className="h-screen w-full flex">
      <aside className="w-1/4 h-full bg-montain bg-center bg-cover shadow-lg shadow-gray-400 relative">
        <div className="w-full h-full bg-slate-900 bg-opacity-50 backdrop-blur-sm p-5 flex flex-col items-center gap-16">
          <div className="z-10 relative w-1/2">
            <img src={logo} alt="logo_da_empresa" />
          </div>

          <div className="w-full h-2/3 flex flex-col gap-3 px-1">
            {btnList.map((value, key) => (
              <Button 
                icon={value.icon}
                title={value.title}
                active={value.active}
                setFunc={handleAddSet}
                setList={btnSetStatusList.current}
                page={value.page}
                setPage={handleChangePage}
                key={key}
              />
            ))}
          </div>

          <div className="absolute bottom-5 left-5 w-1/3 py-1 text-center font-bold bg-red-500 text-white rounded-md cursor-pointer hover:scale-105">Sair</div>
        </div>
      </aside>
      {actualPage || <Board />}
    </main>
  )
}

export default Dashboard