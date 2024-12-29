import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { SignUp } from './pages/users/SignUp'


function App() {
  return (
    <BrowserRouter>
      <Routes>
        //todo: パス管理を専用にする
        <Route path="/signup" element={<SignUp />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
