import { Field, InputType } from "@nestjs/graphql";
import { IsNotEmpty, IsOptional, IsString, IsNumber } from "class-validator";
import { FindAllOrderRequest } from "__generated__/proto/order";

@InputType()
export class FindAllOrderInput implements FindAllOrderRequest {
  @Field(() => Number)
  @IsNotEmpty()
  @IsNumber()
  readonly limit?: number;

  @Field(() => Number)
  @IsNotEmpty()
  @IsNumber()
  readonly offset?: number;

  @Field(() => String)
  @IsNotEmpty()
  readonly orderBy?: ORDER_BY_TYPE;

  @Field(() => String, { nullable: true })
  @IsOptional()
  readonly orderType?: ORDER_TYPE_TYPE;
}
