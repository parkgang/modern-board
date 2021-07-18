import { useState, cloneElement, SyntheticEvent } from "react";
import { useQueryClient } from "react-query";
import { useErrorHandler } from "react-error-boundary";
import { Button, Form, FormInput, Flex } from "@fluentui/react-northstar";

import { ConfirmDialog } from "components/Dialog";

import { PutUser } from "libs/api/user";

import { WrapError } from "libs/error";

import { UserDTO } from "types/user";

type TUserFormTarget = {
  [K in keyof UserDTO]: { value: string };
};

type Props = {
  id: number;
  trigger: JSX.Element;
};

export default function UserForm({ id, trigger }: Props) {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const handleError = useErrorHandler();
  const queryClient = useQueryClient();

  function handlerTrigger() {
    setIsOpen(true);
  }
  async function handlerSubmit(e: SyntheticEvent) {
    try {
      e.preventDefault();
      const target = e.target as typeof e.target & TUserFormTarget;
      const name = target.name.value;
      const arg = target.arg.value;
      await PutUser(id, {
        name: name,
        arg: parseInt(arg),
      });
      queryClient.invalidateQueries("userList");
      setIsOpen(false);
    } catch (error) {
      WrapError(error, `handlerDelete 에러`);
      handleError(error);
    }
  }
  function handlerCancel() {
    setIsOpen(false);
  }

  return (
    <>
      {cloneElement(trigger, {
        onClick: handlerTrigger,
      })}
      <ConfirmDialog
        isOpen={isOpen}
        content={
          <>
            <Form onSubmit={handlerSubmit}>
              <FormInput label="이름" name="name" required type="text" fluid />
              <FormInput label="나이" name="arg" required type="number" fluid />
              <Flex gap="gap.small" hAlign="center">
                <Button content="생성" primary />
                <Button content="취소" onClick={handlerCancel} />
              </Flex>
            </Form>
          </>
        }
      />
    </>
  );
}
