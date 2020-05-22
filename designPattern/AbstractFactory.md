# 抽象工厂模式 (Abstract Factory)
---

## 1.1.1 目的

在不指定具体类的情况下创建一系列相关或依赖对象。通常创建的类都实现相同的接口。抽象工厂的客户并不关心这些对象是如何创建的，它只是知道他们是如何一起运行的。

## 1.1.2 UML图
![抽象工厂模式UML图](https://cdn.learnku.com/uploads/images/201905/27/1/BMhC0s5JAh.png!large "抽象工厂模式UML图")

## 1.1.3 PHP代码示例

```php
<?php
interface Product
{
    public function calculatePrice(): int;
}

class ShippableProduct implements Product
{
    /**
     * @var float
     */
    private $productPrice;
    /**
     * @var float
     */
    private $shippingCosts;
    public function __construct(int $productPrice, int $shippingCosts)
    {
        $this->productPrice = $productPrice;
        $this->shippingCosts = $shippingCosts;
    }
    public function calculatePrice(): int
    {
        return $this->productPrice + $this->shippingCosts;
    }
}

class DigitalProduct implements Product
{
    /**
     * @var int
     */
    private $price;
    public function __construct(int $price)
    {
        $this->price = $price;
    }
    public function calculatePrice(): int
    {
        return $this->price;
    }
}
```
