import { Box, Card, Center, Input, Stack, } from "@chakra-ui/react"
import { BaseLayout } from "@/pages/layout/BaseLayout"
import { Button } from "@/components/ui/button"
import { useForm } from "react-hook-form"
import { Helper } from "@/libs/helper"
import { Field } from "@/components/ui/field"
import { CreateAccountRequest } from "@/internal/ports/impls/accounts"
import { AccountApi } from "@/internal/apis/account-api"
import { Toaster, toaster } from "@/components/ui/toaster"
import { useState } from "react"

export const SignUp = () => {
    const [isLoading, setIsLoading] = useState(false)
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<CreateAccountRequest>()
    const api = new AccountApi();
    const onSubmit = handleSubmit(async (data) => {
        setIsLoading(true)
        const req = new CreateAccountRequest(data)
        const res = await api.signup(req)
        setIsLoading(false)
        if (!res.isSuccess()) {
            if (res.isEmailAlreadyExistsError()) {
                toaster.create({
                    description: "このメールアドレスは既に登録されています",
                    type: "error",
                })
                return
            }
            if (res.isPasswordInvalidError()) {
                toaster.create({
                    description: "パスワードが不正です",
                    type: "error",
                })
                return
            }
            toaster.create({
                description: JSON.stringify(res.getData()),
                type: "error",
            })
            return
        }
        console.log(data)
        console.log(res)
    })
    const log = () => {
        console.log("clicked")
        console.log()
    }
    return (
        <BaseLayout>
            <Toaster />
            <Center>
                <Box data-state="open" _open={{ animation: "slide-from-right-full 500ms ease-out" }}>
                    <Card.Root maxW="sm" variant={"elevated"} flexDirection="column" alignItems="center" m="20">
                        <Card.Header>
                            <Card.Title>Sign up</Card.Title>
                            <Card.Description>
                                Fill in the form below to create an account
                            </Card.Description>
                        </Card.Header>
                        <Card.Body w="full">
                            <form onSubmit={onSubmit}>
                                <Stack color="GrayText" gap="4" w="full">
                                    <Field
                                        label="メールアドレス"
                                        invalid={!!errors.email}
                                        errorText={errors.email?.message}
                                    >
                                        <Input {...register("email", { required: "メールアドレスを入力してください", pattern: Helper.getRegExpEmail() })} placeholder="mail@example.com" className="peer" />
                                    </Field>
                                    <Field
                                        label="ユーザーネーム"
                                        invalid={!!errors.name}
                                        errorText={errors.name?.message}
                                    >
                                        <Input {...register("name", { required: "ユーザーネームを入力してください", min: 3 })} placeholder="Name" />
                                    </Field>
                                    <Field
                                        label="パスワード"
                                        invalid={!!errors.password}
                                        errorText={errors.password?.message}
                                        style={{ whiteSpace: "pre-line" }}
                                    >
                                        <Input {...register("password", { required: "パスワードを入力してください", pattern: { value: Helper.getRegExpPassword(), message: "・英数字混合10文字以上\n・大文字小文字混合\n・特殊記号１文字以上" } })} placeholder="Password" />
                                    </Field>
                                    <Button type="submit" mt="2rem" onClick={log}
                                        loading={isLoading}
                                        loadingText="登録"
                                    >登録</Button>
                                </Stack>
                            </form>
                        </Card.Body>
                    </Card.Root>
                </Box>
            </Center>
        </BaseLayout>
    )
}

// const floatingStyles = defineStyle({
//     pos: "absolute",
//     bg: "bg",
//     px: "0.5",
//     top: "-3",
//     insetStart: "2",
//     fontWeight: "normal",
//     pointerEvents: "none",
//     transition: "position",
//     _peerPlaceholderShown: {
//         color: "fg.muted",
//         top: "2.5",
//         insetStart: "3",
//     },
//     _peerFocusVisible: {
//         color: "fg",
//         top: "-3",
//         insetStart: "2",
//     },
// })