<?php

declare(strict_types=1);

class HighScores
{
    public $scores = null;
    public $latest = null;
    public $personalBest = null;
    public $personalTopThree = [];

    public function __construct(array $scores)
    {
        $count = count($scores);
        $this->scores = $scores;

        rsort($this->scores);

        if($count > 0){
            $this->latest = $scores[($count - 1)];
            $this->personalBest = $this->scores[0];
            $this->personalTopThree = array_slice($this->scores, 0, 3);
        }

        $this->scores = $scores;
    }
}
