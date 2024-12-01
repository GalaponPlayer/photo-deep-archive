import { Header } from "@/components/foundation/Header"
import { Box } from "@chakra-ui/react"

export const BaseLayout = ({ children }) => {
    return (
        <div>
            <Header />
            <main>
                <Box bg="gray.400" minH="100vh">{children}</Box>
            </main>
        </div>
    )
}