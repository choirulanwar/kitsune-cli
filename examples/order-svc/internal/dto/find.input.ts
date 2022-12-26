import { Field, InputType } from "@nestjs/graphql";
import { IsNotEmpty, IsOptional, IsString, IsNumber } from "class-validator";
import { FindOrderRequest } from "__generated__/proto/order";

@InputType()
export class FindOrderInput implements FindOrderRequest {
  @Field(() => String)
  @IsNotEmpty()
  @IsString()
  readonly id?: string;
}
