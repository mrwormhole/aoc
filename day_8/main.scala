import scala.io.Source
import scala.collection.mutable.ArrayBuffer
import scala.util.control.Breaks._
import scala.compiletime.ops.int

object Main {
    def main(args: Array[String]) = {
        val filename: String = "input.txt"
        
        var points: ArrayBuffer[ArrayBuffer[Int]] = ArrayBuffer()
        for (line <- Source.fromFile(filename).getLines()) {
            var rowPoints: ArrayBuffer[Int] = ArrayBuffer()
            for (n <- line) {
                rowPoints.append(n.toString().toInt)
            }
            points.append(rowPoints)
        }

        def isDirectionVisible(value: Int, direction: ArrayBuffer[Int]): Boolean = {
            if (direction.length == 1 && value <= direction(0)) {
                return false
            }
        
            var visible = true
            breakable{
                for (i <- 0 to direction.length - 1) {
                    if (value <= direction(i)) {
                        visible = false
                        break
                    }
                }
            }
            return visible
        }

        def calculateScenicScore(value: Int, direction: ArrayBuffer[Int]): Int = {
            var result: Int = 1
            if (direction.length == 1) {
                return result
            }

            for (i <- 0 to direction.length - 2) {
                if (value <= direction(i)) {
                    return result
                }
                result = result + 1
            }
            return result
        }

        var sum: Int = (points.length - 2) * 2 + (points(0).length) * 2
        var bestScenicScore: Int = 0
        for(i <- 1 to points.length-2; j <- 1 to points(i).length-2) {
            var shouldAdd: Boolean = false
            val left: ArrayBuffer[Int] = points.clone()(i).take(j).reverse
            if (isDirectionVisible(points(i)(j), left)) {
                shouldAdd = true
            }
            val right: ArrayBuffer[Int] = points.clone()(i).takeRight(points(i).length - j - 1)
            if (isDirectionVisible(points(i)(j), right)) {
                shouldAdd = true
            }
            var transposed: ArrayBuffer[ArrayBuffer[Int]] = points.clone().transpose
            val top: ArrayBuffer[Int] = transposed(j).take(i).reverse
            if (isDirectionVisible(points(i)(j), top)) {
                shouldAdd = true
            }
            transposed = points.clone().transpose
            val bottom: ArrayBuffer[Int] = transposed(j).takeRight(points.length - i - 1)
            if (isDirectionVisible(points(i)(j), bottom)) {
                shouldAdd = true
            }
            if (shouldAdd) {
                var scenicScore: Int = 1
                scenicScore *= calculateScenicScore(points(i)(j), left)
                scenicScore *= calculateScenicScore(points(i)(j), right)
                scenicScore *= calculateScenicScore(points(i)(j), top)
                scenicScore *= calculateScenicScore(points(i)(j), bottom)
                sum += 1
                if (scenicScore > bestScenicScore) {
                    bestScenicScore = scenicScore
                }
            }    
        }
        println(s"part 1: ${sum}")
        println(s"part 2: ${bestScenicScore}")
    }
}