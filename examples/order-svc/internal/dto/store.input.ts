import { Field, InputType } from "@nestjs/graphql";
import { IsNotEmpty, IsOptional, IsString, IsNumber } from "class-validator";
import { StoreOrderRequest } from "__generated__/proto/order";

@InputType()
export class StoreOrderInput implements StoreOrderRequest {
  @Field(() => String)
  @IsNotEmpty()
  readonly type?: PRODUCT_ORDER_TYPE;

  @Field(() => [ProductItemOrderInput])
  @IsNotEmpty()
  readonly products?: ProductItemOrderInput[];

  @Field(() => String, { nullable: true })
  @IsOptional()
  @IsString()
  readonly promo_id?: string;

  @Field(() => StoreAddressInput, { nullable: true })
  @IsOptional()
  readonly delivery_location?: StoreAddressInput;

  @Field(() => String)
  @IsNotEmpty()
  readonly payment_method?: PAYMENT_METHOD_TYPE;
}
