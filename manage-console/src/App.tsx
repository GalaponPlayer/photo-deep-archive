import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { SignUp } from './pages/users/SignUp'


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signup" element={<SignUp />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
