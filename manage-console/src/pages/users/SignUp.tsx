import { Box, Card, Container, Heading, Input, Stack } from "@chakra-ui/react"
import { BaseLayout } from "@/pages/layout/BaseLayout"
import { Button } from "@/components/ui/button"
import { useForm } from "react-hook-form"
import { Helper } from "@/libs/helper"
import { Field } from "@/components/ui/field"
import { CreateAccountRequest } from "@/internal/ports/impls/accounts"

export const SignUp = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<CreateAccountRequest>()
    const onSubmit = handleSubmit((data) => {
        console.log(data)
    })
    return (
        <BaseLayout>
            <Container centerContent>
                <Box data-state="open" _open={{ animation: "slide-from-right-full 500ms ease-out" }}>
                    <Card.Root w="352px" bg="blue.100" border="none" shadow="inset" flexDirection="column" alignItems="center">
                        <Card.Body>
                            <form onSubmit={onSubmit}>
                                <Stack color="GrayText">
                                    <Heading as="h2" size="md">Sign Up</Heading>
                                    <Field
                                        label="Email"
                                        invalid={!!errors.email}
                                        errorText={errors.email?.message}
                                    >
                                        <Input {...register("email", { required: "メールアドレスを入力してください", pattern: Helper.getRegExpEmail() })} placeholder="Email" />
                                    </Field>
                                    <Field
                                        label="Password"
                                        invalid={!!errors.password}
                                        errorText={errors.password?.message}
                                    >
                                        <Input {...register("password", { required: "・数字１文字以上\n・特殊記号１文字以上\n大文字小文字をそれぞれ１文字以上", pattern: Helper.getRegExpPassword() })} placeholder="Password" />
                                    </Field>
                                    <Button type="submit" mt="2rem">登録</Button>
                                </Stack>
                            </form>
                        </Card.Body>
                    </Card.Root>
                </Box>
            </Container>
        </BaseLayout>
    )
}
