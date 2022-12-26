import { Field, InputType } from "@nestjs/graphql";
import { IsNotEmpty, IsOptional, IsString, IsNumber } from "class-validator";
import { DeleteOrderRequest } from "__generated__/proto/order";

@InputType()
export class DeleteOrderInput implements DeleteOrderRequest {
  @Field(() => String)
  @IsNotEmpty()
  @IsString()
  readonly id?: string;
}
