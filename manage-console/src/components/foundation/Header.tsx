import { Box, Flex, Heading, Link } from "@chakra-ui/react"
import { Avatar } from "@/components/ui/avatar"

export const Header = () => {
    return (
        <header>
            <Box bg="cyan.950" p="4" color="cyan.50">
                <Flex justify="space-between" align="center">
                    <Heading>Photo-Deep-Archive</Heading>
                    <Flex>
                        <Link href="/signup">Sign Up</Link>
                        <Avatar name="User" />
                    </Flex>
                </Flex>
            </Box>
        </header>
    )
}